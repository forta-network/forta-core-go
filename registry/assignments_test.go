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
		AgentID:          "0x492c05269cbefe3a1686b999912db1fb5a39ce2e4578ac3951b0542440f435d9",
		AgentManifest:    "QmaveKpceDCaersFd5C24DFzAB2UQuHGcz1dJZhqkQiMBn",
		AssignedScanners: 4,
		ScannerIndex:     1,
		AgentOwner:       "0x88dC3a2284FA62e0027d6D6B1fCfDd2141a143b8",
	},
	{
		AgentID:          "0x98b87a29ecb6c8c0f8e6ea83598817ec91e01c15d379f03c7ff781fd1141e502",
		AgentManifest:    "QmaToWbUsVKXGpiqcxEBS5aQYAiXXPtWDqKZ798WJAftYE",
		AssignedScanners: 45,
		ScannerIndex:     39,
		AgentOwner:       "0x1D179DE1E32Cf7aE63d208EE6201222E299298c9",
	},
	{
		AgentID:          "0x0b241032ca430d9c02eaa6a52d217bbff046f0d1b3f3d2aa928e42a97150ec91",
		AgentManifest:    "QmVzykcTd5u73HCn8CwikzboTveoCHTgciCa9whwXSmZx9",
		AssignedScanners: 3,
		ScannerIndex:     0,
		AgentOwner:       "0xB7EC4b33a3228f46901506De71673906646aE678",
	},
	{
		AgentID:          "0x0f21668ebd017888e7ee7dd46e9119bdd2bc7f48dbabc375d96c9b415267534c",
		AgentManifest:    "QmS6wdfta6Z4toqCxisqiVUgUvM8W38h2V5a2qLpreQedi",
		AssignedScanners: 1,
		ScannerIndex:     0,
		AgentOwner:       "0x70358461af09dCb10B003e8e0a6033de34387A2a",
	},
	{
		AgentID:          "0x59413fe6f6be7eb100e40f607fb448ef2a5bb838a5d19e6246fe70e3269ffd56",
		AgentManifest:    "QmaCD22UohViRyYPPh3oZnN8pJZinycbNM7zUHPY9WFPFt",
		AssignedScanners: 5,
		ScannerIndex:     2,
		AgentOwner:       "0x6aea36C4a9Cf8AC053c07E662fa27e8BDF047121",
	},
	{
		AgentID:          "0x1d646c4045189991fdfd24a66b192a294158b839a6ec121d740474bdacb3ab23",
		AgentManifest:    "QmUTNe9vVAUvmLaAfwffesbwCmWB8DigTUgcr8jFnHa5S4",
		AssignedScanners: 48,
		ScannerIndex:     3,
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
	},
	{
		AgentID:          "0x47c45816807d2eac30ba88745bf2778b61bc106bc76411b520a5289495c76db8",
		AgentManifest:    "QmPWyYJpNTSx2GU33ict4VjPc8tojbdAoRHtgHvo3Zdgtj",
		AssignedScanners: 48,
		ScannerIndex:     20,
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
	},
	{
		AgentID:          "0xbb34c89f51bba941ef7ad54daccbc82eadb6177ffd2183d0939561499e295424",
		AgentManifest:    "QmPPDrFo1ZpGaRjjKq8tZ6XZCSHPoYbSUuLh1E5jTqWYK5",
		AssignedScanners: 3,
		ScannerIndex:     0,
		AgentOwner:       "0xBc9ebd3d23Fd48b521b03343a12e31F3D6093F70",
	},
	{
		AgentID:          "0x1a69f5ec8ef436e4093f9ec4ce1a55252b7a9a2d2c386e3f950b79d164bc99e0",
		AgentManifest:    "Qmev376tngMcjs2qg6PNUqNbjh7FUXkGQcnQsod56R8WL9",
		AssignedScanners: 12,
		ScannerIndex:     6,
		AgentOwner:       "0x88dC3a2284FA62e0027d6D6B1fCfDd2141a143b8",
	},
	{
		AgentID:          "0x8fe07f1a4d33b30be2387293f052c273660c829e9a6965cf7e8d485bcb871083",
		AgentManifest:    "bafkreihxf27on2447l42fowryq6qvvy4gtdqeh3dsnygbbobcyajpp3qrm",
		AssignedScanners: 3,
		ScannerIndex:     0,
		AgentOwner:       "0x40e98Fe8BCc8c0fF48A585bc6f28d504fF17A927",
	},
	{
		AgentID:          "0xabc0bb6fe5e0d0b981dec4aa2337ce91676358c6e8bf1fec06cc558f58c3694e",
		AgentManifest:    "QmUPptyURyfvNWwx5CqJP6iKjVD6LCBBSP5sHLiYaNY4vb",
		AssignedScanners: 3,
		ScannerIndex:     2,
		AgentOwner:       "0xF08021332fec194Ff140eF8d5b91868EdFD1A879",
	},
	{
		AgentID:          "0xe7662cdcf453a8bf1f88955d6ce7a585ca107f7174b73efb92be6cfdac6a8f6a",
		AgentManifest:    "QmemAGdw4zv2nzH6zwjwpuJz1oDauJS5LYUPjHAJxXq7Dc",
		AssignedScanners: 3,
		ScannerIndex:     1,
		AgentOwner:       "0x88dC3a2284FA62e0027d6D6B1fCfDd2141a143b8",
	},
	{
		AgentID:          "0x6f022d4a65f397dffd059e269e1c2b5004d822f905674dbf518d968f744c2ede",
		AgentManifest:    "QmQKZeH8Jfb3emicac7dThZBrd9ZiqiYEVFcL74ae4bPw9",
		AssignedScanners: 40,
		ScannerIndex:     34,
		AgentOwner:       "0xAF1Bb3e4A1F8401fE044B51d0eAB41E917b36803",
	},
	{
		AgentID:          "0xb4e4c2584edab51f0c8ed56501d8974d1ba3bb798007a20ebf8a55d5f1410907",
		AgentManifest:    "QmQKqEoxnc4gUhPiVsZ4CJLamshAxCLrBRSKk9ZDB5uxGt",
		AssignedScanners: 3,
		ScannerIndex:     1,
		AgentOwner:       "0xc074A2949112218767B631449957626Ca3a3ff45",
	},
	{
		AgentID:          "0x7cfeb792e705a82e984194e1e8d0e9ac3aa48ad8f6530d3017b1e2114d3519ac",
		AgentManifest:    "QmezfraHfiAwzkt6FKsSn1ZU6st2v8nrcmSwGLgDMQ3ure",
		AssignedScanners: 36,
		ScannerIndex:     10,
		AgentOwner:       "0x88dC3a2284FA62e0027d6D6B1fCfDd2141a143b8",
	},
	{
		AgentID:          "0x36c962b41bdc05788994a24a485e1438ea1a106822fc7fb785c5079128daf3a8",
		AgentManifest:    "bafkreieumpoiosco5etxseuthnfqlnnn6qgajzdxrrxetujwo7epkgwrai",
		AssignedScanners: 3,
		ScannerIndex:     0,
		AgentOwner:       "0x137D5E3D0e6fDD02d5944468C09d9A6FE2D9ef1d",
	},
	{
		AgentID:          "0xb27524b92bf27e6aa499a3a7239232ad425219b400d3c844269f4a657a4adf03",
		AgentManifest:    "QmYhi8pYKZr8D2gDuUGVYYW5u7c9iuJJJJrxZmo1XekfRS",
		AssignedScanners: 48,
		ScannerIndex:     14,
		AgentOwner:       "0x2835a787d8d74724181F97ec97d9882dD7b3F2be",
	},
}

func TestGetAssignmentList(t *testing.T) {
	r := require.New(t)

	cfg := defaultConfig
	cfg.NoRefresh = true
	client, err := NewClient(context.Background(), cfg)
	r.NoError(err)

	scannerAddress := "0x03c80fe4a7ff15aff625d1e25d4969012b6ca8cc"
	assignedChainID := big.NewInt(1)
	blockNumber := big.NewInt(43896826)
	var assignments []*Assignment
	for i := 0; i < 11; i++ {
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
