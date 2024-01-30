package security

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/forta-network/forta-core-go/encoding"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
)

var ErrMissingSignature = errors.New("missing signature")
var ErrInvalidSignature = errors.New("invalid signature")

func ReadPassphrase() (string, error) {
	f, err := os.OpenFile("/passphrase", os.O_RDONLY, 0400)
	if err != nil {
		return "", err
	}
	pw, err := io.ReadAll(bufio.NewReader(f))
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

// LoadKey loads the passphrase and the node private key.
func LoadKey(keysDirPath string) (*keystore.Key, error) {
	passphrase, err := ReadPassphrase()
	if err != nil {
		return nil, err
	}
	return LoadKeyWithPassphrase(keysDirPath, passphrase)
}

// LoadKeyWithPassphrase decrypts and loads the node private key using provided passphrase.
func LoadKeyWithPassphrase(keysDirPath, passphrase string) (*keystore.Key, error) {
	files, err := ioutil.ReadDir(keysDirPath)
	if err != nil {
		return nil, err
	}

	if len(files) != 1 {
		return nil, errors.New("there must be only one key in key directory")
	}

	keyBytes, err := ioutil.ReadFile(path.Join(keysDirPath, files[0].Name()))
	if err != nil {
		return nil, err
	}

	return LoadKeyFromBytes(keyBytes, passphrase)
}

// LoadKeyFromBytes decrypts and loads the node private key using bytes of key and passphrase
func LoadKeyFromBytes(keyBytes []byte, passphrase string) (*keystore.Key, error) {
	return keystore.DecryptKey(keyBytes, passphrase)
}

/*
The alertHash includes
1. alert.Id, which is a hash of all fields for the alert (same for all scanners that find this alert)
2. metadata, which is unique to this scanner (deterministic via list conversion)
3. timestamp, which is unique to this scanner
*/
func alertHash(alert *protocol.Alert) common.Hash {
	metadata := utils.MapToList(alert.Metadata)
	alertStr := fmt.Sprintf("%s%s%s", alert.Id, strings.Join(metadata, ""), alert.Timestamp)
	return crypto.Keccak256Hash([]byte(alertStr))
}

// SignAlert signs the alert using the alertID and deterministicly formatted Metadata
func SignAlert(key *keystore.Key, alert *protocol.Alert) (*protocol.SignedAlert, error) {
	hash := alertHash(alert)
	signature, err := SignBytes(key, hash.Bytes())
	if err != nil {
		return nil, err
	}

	return &protocol.SignedAlert{
		Alert:     alert,
		Signature: signature,
	}, nil
}

// VerifyAlertSignature returns an error if the signature for the signed alert is invalid
func VerifyAlertSignature(sa *protocol.SignedAlert) error {
	if sa.Signature == nil {
		return ErrMissingSignature
	}

	hash := alertHash(sa.Alert)
	return VerifySignature(hash.Bytes(), sa.Signature.Signer, sa.Signature.Signature)
}

func SignBytes(key *keystore.Key, b []byte) (*protocol.Signature, error) {
	hash := crypto.Keccak256(b)
	sig, err := crypto.Sign(hash, key.PrivateKey)

	if err != nil {
		return nil, err
	}

	return &protocol.Signature{
		Signature: fmt.Sprintf("0x%s", hex.EncodeToString(sig)),
		Algorithm: "ECDSA",
		Signer:    key.Address.Hex(),
	}, nil
}

func SignString(key *keystore.Key, input string) (*protocol.Signature, error) {
	return SignBytes(key, []byte(input))
}

func SignerAddressFromSignature(message []byte, sigHex string) (string, error) {
	sigHex = strings.ReplaceAll(sigHex, "0x", "")
	signature, err := hex.DecodeString(sigHex)
	if err != nil {
		return "", err
	}

	hash := crypto.Keccak256(message)
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*pubKey).Hex(), nil
}

func VerifySignature(message []byte, signerAddress string, sigHex string) error {
	sigHex = strings.ReplaceAll(sigHex, "0x", "")
	signature, err := hex.DecodeString(sigHex)
	if err != nil {
		return err
	}

	hash := crypto.Keccak256Hash(message)

	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return err
	}

	if pubKey == nil {
		return errors.New("could not recover address (pub is nil)")
	}

	addr := crypto.PubkeyToAddress(*pubKey)

	if addr.Hex() != signerAddress {
		return ErrInvalidSignature
	}

	return nil
}

func signPayload(key *keystore.Key, payloadType protocol.SignedPayload_PayloadType, msg proto.Message) (*protocol.SignedPayload, error) {
	encoded, err := encoding.EncodeGzippedProto(msg)
	if err != nil {
		return nil, err
	}

	signature, err := SignString(key, encoded)
	if err != nil {
		return nil, err
	}

	return &protocol.SignedPayload{
		Type:      payloadType,
		Encoded:   encoded,
		Signature: signature,
	}, nil

}

// SignBatch will sign an alert batch and return a SignedAlertBatch
func SignBatch(key *keystore.Key, payload *protocol.AlertBatch) (*protocol.SignedPayload, error) {
	return signPayload(key, protocol.SignedPayload_BATCH, payload)
}

// SignBatchSummary will sign an alert batch summary
func SignBatchSummary(key *keystore.Key, payload *protocol.BatchSummary) (*protocol.SignedPayload, error) {
	return signPayload(key, protocol.SignedPayload_BATCH_SUMMARY, payload)
}

// SignBatchReceipt will sign a batch receipt
func SignBatchReceipt(key *keystore.Key, payload *protocol.BatchReceipt) (*protocol.SignedPayload, error) {
	return signPayload(key, protocol.SignedPayload_BATCH_RECEIPT, payload)
}

// VerifySignedPayload will return an error if the signature fails to validate
func VerifySignedPayload(payload *protocol.SignedPayload) error {
	if payload.Signature == nil {
		return ErrMissingSignature
	}
	return VerifySignature([]byte(payload.Encoded), payload.Signature.Signer, payload.Signature.Signature)
}

// NewTransactOpts creates new opts with the private key.
func NewTransactOpts(key *keystore.Key) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key.PrivateKey)
}

// DecodeEthereumSignature decodes an Ethereum signature.
func DecodeEthereumSignature(sigHex string) ([]byte, error) {
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return nil, err
	}

	// logic from https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return nil, errors.New("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	return sig, nil
}

// EncodeEthereumSignature encodes an Ethereum signature.
func EncodeEthereumSignature(sig []byte) (string, error) {
	// logic from https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 0 && sig[64] != 1 {
		return "", errors.New("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] += 27

	return hexutil.Encode(sig), nil
}
