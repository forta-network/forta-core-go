package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/forta-network/forta-core-go/registry"
	log "github.com/sirupsen/logrus"
)

type scannerSummary struct {
	Scanner  *registry.Scanner `json:"scanner"`
	AgentIDs []string          `json:"agentIds"`
}

type agentSummary struct {
	Agent      *registry.Agent `json:"agent"`
	ScannerIDs []string        `json:"scannerIds"`
}

type export struct {
	Agents   []*agentSummary   `json:"agents"`
	Scanners []*scannerSummary `json:"scanners"`
}

func main() {
	ex := &export{}
	c, err := registry.NewClient(context.Background(), registry.ClientConfig{
		JsonRpcUrl: os.Getenv("POLYGON_JSON_RPC"), //TODO: fill in your json rpc provider
		ENSAddress: "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7",
		Name:       "export",
	})
	if err != nil {
		log.WithError(err).Fatal("failed to initialize client")
		return
	}
	if err := c.PegLatestBlock(); err != nil {
		log.WithError(err).Fatal("failed to peg block for client")
		return
	}

	err = c.ForEachScanner(func(s *registry.Scanner) error {
		scn := &scannerSummary{
			Scanner:  s,
			AgentIDs: []string{},
		}
		ex.Scanners = append(ex.Scanners, scn)
		return c.ForEachAssignedAgent(s.ScannerID, func(a *registry.Agent) error {
			scn.AgentIDs = append(scn.AgentIDs, a.AgentID)
			return nil
		})
	})
	if err != nil {
		log.WithError(err).Fatal("failed to get scanners for client")
		return
	}

	err = c.ForEachAgent(func(a *registry.Agent) error {
		agt := &agentSummary{
			Agent:      a,
			ScannerIDs: []string{},
		}
		ex.Agents = append(ex.Agents, agt)
		return c.ForEachAssignedScanner(a.AgentID, func(s *registry.Scanner) error {
			agt.ScannerIDs = append(agt.ScannerIDs, s.ScannerID)
			return nil
		})
	})
	if err != nil {
		log.WithError(err).Fatal("failed to get scanners for client")
		return
	}

	b, err := json.Marshal(ex)
	if err != nil {
		log.WithError(err).Fatal("failed to marshal export json")
		return
	}

	fmt.Println(string(b))
}
