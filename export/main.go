package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/forta-protocol/forta-core-go/registry"
	log "github.com/sirupsen/logrus"
)

type scannerSummary struct {
	Scanner *registry.Scanner `json:"scanner"`
	Agents  []*registry.Agent `json:"agents"`
}

type agentSummary struct {
	Agent    *registry.Agent     `json:"agent"`
	Scanners []*registry.Scanner `json:"scanners"`
}

type export struct {
	Agents   []*agentSummary
	Scanners []*scannerSummary
}

func main() {
	ex := &export{}
	c, err := registry.NewClient(context.Background(), registry.ClientConfig{
		JsonRpcUrl: "", //TODO: fill in your json rpc provider
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
			Scanner: s,
		}
		ex.Scanners = append(ex.Scanners, scn)
		return c.ForEachAssignedAgent(s.ScannerID, func(a *registry.Agent) error {
			scn.Agents = append(scn.Agents, a)
			return nil
		})
	})
	if err != nil {
		log.WithError(err).Fatal("failed to get scanners for client")
		return
	}

	err = c.ForEachAgent(func(a *registry.Agent) error {
		agt := &agentSummary{
			Agent: a,
		}
		ex.Agents = append(ex.Agents, agt)
		return c.ForEachAssignedScanner(a.AgentID, func(s *registry.Scanner) error {
			agt.Scanners = append(agt.Scanners, s)
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
