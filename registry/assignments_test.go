package registry

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var testExpectedAssignmentList = []*Assignment{
	{
		AgentID:          "0x441d3228a68bbbcf04e6813f52306efcaf1e66f275d682e62499f44905215250",
		AgentManifest:    "Qmagb2oHhcRD9orpCz1xbnUCC53QKK33UgkjZaA7Vdsxfx",
		AgentOwner:       "0x88dC3a2284FA62e0027d6D6B1fCfDd2141a143b8",
		AssignedScanners: 36,
		ScannerIndex:     11,
	},
	{
		AgentID:          "0x47c45816807d2eac30ba88745bf2778b61bc106bc76411b520a5289495c76db8",
		AgentManifest:    "QmZRnE2uBavyaD8BdARhQ32C2x3VbXLfCqiVY9p3jpJYyA",
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
		AssignedScanners: 47,
		ScannerIndex:     3,
	},
	{
		AgentID:          "0x1d5aa6e0085208814ac479fb389fe1295a35ef698c2543eaf93ca27e2e5a84a1",
		AgentManifest:    "QmbVUJjFm4VpwyEYAGRwmDXVJKDFa3tX4mdi2wwtNTxioR",
		AgentOwner:       "0x88dC3a2284FA62e0027d6D6B1fCfDd2141a143b8",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0x60cfaeedd7561f2f2cb6f4a508be32598b877c715646c3f175b6a3eac46f94e9",
		AgentManifest:    "QmYwybxFG8rrLzkPXUtpNPcneh5tu31sTHNEvvL2cFbSMV",
		AgentOwner:       "0x81813Bd41459c0Bb757308C855c0521A4e308E36",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0xe52aecd3413346bb4950f5798e304b76d697c7782bb312eab451a097c78163b9",
		AgentManifest:    "bafkreie57puqop2cgpns6v5ycb65xakkm3fk6kdcmaxtmkphntwtxwug2i",
		AgentOwner:       "0x70358461af09dCb10B003e8e0a6033de34387A2a",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0xe97dd4bb604c044fe0eed9dc3a3450da7907aae63f20378ca5cf1a96750bed76",
		AgentManifest:    "bafkreieh2c456fmkaumagkrl6tqn6hqo4dajgxfqemcrnaedh4cjjjwzse",
		AgentOwner:       "0x0E860F44d73F9FDbaF5E9B19aFC554Bf3C8E8A57",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0x92cec8bacc8436cc59becf2abc094047a150224a626fd506fa8d64c0925bdd61",
		AgentManifest:    "Qmey7vqiDDtygSoDNf26ZLH6wXgbhVBTjZ7gU5MQrb4JoL",
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0x34c27c43e0a45bced8f8a941b3d552f5e6feae62afd7e2e88b5024f7de5a8ba0",
		AgentManifest:    "QmeX3TJ2nk6g5SiVFVc3oCiv38WcfzXWejEkwNDH98NBpL",
		AgentOwner:       "0x51F3B9d02e0C2EAB8f4EbC0Ee2E571233083bAec",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0x1d646c4045189991fdfd24a66b192a294158b839a6ec121d740474bdacb3ab23",
		AgentManifest:    "QmUTNe9vVAUvmLaAfwffesbwCmWB8DigTUgcr8jFnHa5S4",
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
		AssignedScanners: 48,
		ScannerIndex:     37,
	},
	{
		AgentID:          "0xa17bd7957d25bc4508e390921a4e54c6475d5271f824ae19734d376db2699556",
		AgentManifest:    "QmWE4fdfPzn7kMbgHZZnsDKzUSPd9Jf2XZ7MLXCh3unhEL",
		AgentOwner:       "0x8bCCFd42Afd3CA58c9c705b368a1Df8Dc7f86E4E",
		AssignedScanners: 3,
		ScannerIndex:     2,
	},
	{
		AgentID:          "0xfb3c0aa7c428c7d4e791b756910a29f6e99e3d8ba4fbf01eb0c8bcfc704937e0",
		AgentManifest:    "QmemMMJmHxU91SxxjWj52LT9TdYpR4aJgtE6cq9bmUBY8L",
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
		AssignedScanners: 48,
		ScannerIndex:     47,
	},
}

func TestGetAssignmentList(t *testing.T) {
	r := require.New(t)

	cfg := defaultConfig
	cfg.NoRefresh = true
	client, err := NewClient(context.Background(), cfg)
	r.NoError(err)

	scannerAddress := "0x5fc518115f29cf5a2424bc9efb9a02acfd77f0a8"
	assignedChainID := big.NewInt(1)
	blockNumber := big.NewInt(43643503)

	var assignments []*Assignment
	for i := 0; i < 10; i++ {
		start := time.Now()
		assignments, err = client.GetAssignmentList(blockNumber, assignedChainID, scannerAddress)
		t.Log("took", time.Since(start))
		if err == nil {
			break
		}
		t.Logf("failed to get assignments - retrying: %v", err)
		time.Sleep(time.Second * 2)
	}

	r.Equal(testExpectedAssignmentList, assignments, "failed to load assignments")
}
