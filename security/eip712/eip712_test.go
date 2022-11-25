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
	testExpectedSigHex       = "0x6898c0bb48618c8edd3a8016b214a66fb6eb5ae096dbc1b8e305919a3b90f788313711a3246b83211704ffba7105088975eb9b86709a493828ca30c6c25459ac1c"
)

func TestSignScannerRegistration(t *testing.T) {
	r := require.New(t)

	scannerKey, err := crypto.HexToECDSA(testScannerPrivateKeyHex)
	r.NoError(err)

	ts, err := time.Parse(time.RFC3339, "2022-11-07T00:00:00Z")
	r.NoError(err)

	verifyingContractChainID := big.NewInt(10)
	scannerChainID := big.NewInt(20)

	encodedData, sig, err := SignScannerRegistration(scannerKey, testVerifyingContract, verifyingContractChainID, &ScannerNodeRegistration{
		Scanner:       testScannerAddress,
		ScannerPoolId: big.NewInt(123),
		ChainId:       scannerChainID,
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
