package registry

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/forta-network/go-multicall"
)

func (c *client) GetAssignmentList(blockNumber, assignedChainID *big.Int, scannerAddress string) ([]*Assignment, error) {
	var (
		opts *bind.CallOpts
		err  error
	)

	// set fallback call options
	if blockNumber == nil {
		opts, err = c.getOpts()
		if err != nil {
			return nil, err
		}
	} else {
		opts = c.getBlockOpts(blockNumber)
	}

	dispatch := c.Contracts().Dispatch
	dispatchMulti := c.Contracts().DispatchMulti
	scannerID := utils.ScannerIDHexToBigInt(scannerAddress)

	// 1st eth_call: get the number of bots assigned to this scanner
	agentCount, err := dispatch.NumAgentsFor(opts, scannerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get the agent count: %v", err)
	}

	// 2nd eth_call: get all bots by iterating over the list
	var agentCalls []*multicall.Call
	for i := int64(0); i < agentCount.Int64(); i++ {
		agentCalls = append(
			agentCalls, dispatchMulti.NewCall(new(agentRefAtOutput), "agentRefAt", scannerID, big.NewInt(i)),
		)
	}

	_, err = c.multiCaller.CallChunked(opts, 100, agentCalls...)
	if err != nil {
		return nil, fmt.Errorf("agentRefAt calls failed: %v", err)
	}

	// 3rd eth_call: for each assigned bot, get the assigned scanner counts
	var numScannersCalls []*multicall.Call
	for i := 0; i < int(agentCount.Int64()); i++ {
		numScannersCall := dispatchMulti.NewCall(
			new(numOutput), "numScannersFor", agentCalls[i].Outputs.(*agentRefAtOutput).AgentId,
		)
		numScannersCalls = append(numScannersCalls, numScannersCall)
	}
	_, err = c.multiCaller.Call(opts, numScannersCalls...)
	if err != nil {
		return nil, fmt.Errorf("numScannersFor calls failed: %v", err)
	}

	// 4th eth_call: get all assigned scanners to all assigned bots
	var scannerCalls []*multicall.Call
	for agentIndex, numScannersCall := range numScannersCalls {
		agent := agentCalls[agentIndex].Outputs.(*agentRefAtOutput)
		scannerCount := numScannersCall.Outputs.(*numOutput).Num.Int64()
		for i := int64(0); i < scannerCount; i++ {
			scannerCall := dispatchMulti.NewCall(
				new(scannerRefAtOutput), "scannerRefAt", agent.AgentId, big.NewInt(i),
			)
			scannerCalls = append(scannerCalls, scannerCall)
		}
	}
	// the amount of scanners can scale up unexpectedly sometimes so this
	// chunking is a protection against that
	_, err = c.multiCaller.CallChunked(opts, 100, scannerCalls...)
	if err != nil {
		return nil, fmt.Errorf("scannerRefAt calls failed: %v", err)
	}

	// construct each bot assignment from what we have collected:
	//  - bot id
	//  - bot manifest
	//  - assigned scanners to the bot, for given chain id
	//  - index of the scanner among assigned, for given chain id: we look at the portion from the previous step
	var (
		assignments     []*Assignment
		scannersChecked int
	)
	for i := 0; i < int(agentCount.Int64()); i++ {
		agent := agentCalls[i].Outputs.(*agentRefAtOutput)
		scannerCount := int(numScannersCalls[i].Outputs.(*numOutput).Num.Int64())

		// the portion of scanners from previous step
		scannersStart := scannersChecked
		scannersEnd := scannersStart + scannerCount
		agentScannersCalls := scannerCalls[scannersStart:scannersEnd]

		var (
			assignedScanners int
			scannerIndex     int
		)
		for _, agentScannersCall := range agentScannersCalls {
			scanner := agentScannersCall.Outputs.(*scannerRefAtOutput)
			// assignedScanners can be treated as the index before it is incremented
			if scanner.ScannerId.Cmp(scannerID) == 0 {
				scannerIndex = assignedScanners
			}
			// count scanners that are only assigned for the given chain
			if scanner.ChainId.Cmp(assignedChainID) == 0 {
				assignedScanners++
			}
		}
		assignments = append(assignments, &Assignment{
			AgentID:          utils.AgentBigIntToHex(agent.AgentId),
			AgentManifest:    agent.Metadata,
			AgentOwner:       agent.Owner.Hex(),
			AssignedScanners: assignedScanners,
			ScannerIndex:     scannerIndex,
		})

		scannersChecked += int(scannerCount) // move to start of the next portion
	}

	return assignments, nil
}
