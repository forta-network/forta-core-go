package eip712

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forta-network/forta-core-go/security"
	"github.com/stretchr/testify/require"
)

var (
	testScannerPrivateKeyHex = "e0b91bdb03d26f15844adf6bb0d51459dbef4bc250caccfa52ab253d5ace32ef"
	testScannerAddress       = common.HexToAddress("0xE980E0A0037f806fC061eC5494FeFE64b9aA9De6")
	testVerifyingContract    = common.HexToAddress("0x83EDD5007f8Bc879B2c5332b1a0Bf724114C7e37")
	testExpectedSigHex       = "0x994087fa0d9e8c8bdac138503cda2d7e4846cf8b92d1e57a9e98977094d2557b3cd1bbad36d84effd3f084b9dc003a268dc83a5a5d4acfabb8d959e0376c0f651c"
)

func TestSignScannerRegistration(t *testing.T) {
	r := require.New(t)

	scannerKey, err := crypto.HexToECDSA(testScannerPrivateKeyHex)
	r.NoError(err)

	ts, err := time.Parse(time.RFC3339, "2022-11-07T00:00:00Z")
	r.NoError(err)

	encodedData, sig, err := SignScannerRegistration(scannerKey, testVerifyingContract, &ScannerNodeRegistration{
		Scanner:       testScannerAddress,
		ScannerPoolId: big.NewInt(123),
		ChainId:       big.NewInt(1),
		Metadata:      "",
		Timestamp:     big.NewInt(ts.Unix()),
	})
	r.NoError(err)
	r.NotNil(encodedData)
	r.NotNil(sig)

	encoded, err := security.EncodeEthereumSignature(sig)
	r.NoError(err)
	r.Equal(testExpectedSigHex, encoded)
}
