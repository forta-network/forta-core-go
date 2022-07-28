package domain

import (
	"bytes"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/goccy/go-json"
	"github.com/golang/protobuf/jsonpb"
	log "github.com/sirupsen/logrus"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
)

type EventType string

const (
	EventTypeBlock EventType = "block"
)

const TimeTrackingTimestampFormat = time.RFC3339Nano

type TrackingTimestamps struct {
	Block       time.Time
	Feed        time.Time
	BotRequest  time.Time
	BotResponse time.Time
}

func (tt *TrackingTimestamps) ToMessage() *protocol.TrackingTimestamps {
	return &protocol.TrackingTimestamps{
		Block:       tt.Block.Format(TimeTrackingTimestampFormat),
		Feed:        tt.Feed.Format(TimeTrackingTimestampFormat),
		BotRequest:  tt.BotRequest.Format(TimeTrackingTimestampFormat),
		BotResponse: tt.BotResponse.Format(TimeTrackingTimestampFormat),
	}
}

func TimeFromString(dt string) time.Time {
	res, err := time.Parse(TimeTrackingTimestampFormat, dt)
	if err != nil {
		log.WithFields(log.Fields{
			"input": dt,
		}).Errorf("invalid time - for safety, ignorning")
	}
	return res
}

func TrackingTimestampsFromMessage(tt *protocol.TrackingTimestamps) *TrackingTimestamps {
	if tt == nil {
		return &TrackingTimestamps{}
	}
	return &TrackingTimestamps{
		Block:       TimeFromString(tt.Block),
		Feed:        TimeFromString(tt.Feed),
		BotRequest:  TimeFromString(tt.BotRequest),
		BotResponse: TimeFromString(tt.BotResponse),
	}
}

type BlockEvent struct {
	EventType  EventType
	ChainID    *big.Int
	Block      *Block
	Logs       []LogEntry
	Traces     []Trace
	Timestamps *TrackingTimestamps
}

func str(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func strArr(vals []*string) []string {
	result := make([]string, 0, len(vals))
	for _, v := range vals {
		result = append(result, str(v))
	}
	return result
}

func strPtr(val string) *string {
	return &val
}

func (t *BlockEvent) ToMessage() (*protocol.BlockEvent, error) {
	evtType := protocol.BlockEvent_BLOCK

	txs := make([]string, 0, len(t.Block.Transactions))
	for _, tx := range t.Block.Transactions {
		txs = append(txs, tx.Hash)
	}
	return &protocol.BlockEvent{
		Type:        evtType,
		BlockHash:   t.Block.Hash,
		BlockNumber: t.Block.Number,
		Network: &protocol.BlockEvent_Network{
			ChainId: utils.BigIntToHex(t.ChainID),
		},
		Block: &protocol.BlockEvent_EthBlock{
			Difficulty:       str(t.Block.Difficulty),
			Hash:             t.Block.Hash,
			Number:           t.Block.Number,
			ParentHash:       t.Block.ParentHash,
			Timestamp:        t.Block.Timestamp,
			Nonce:            str(t.Block.Nonce),
			ExtraData:        str(t.Block.ExtraData),
			GasLimit:         str(t.Block.GasLimit),
			GasUsed:          str(t.Block.GasUsed),
			LogsBloom:        str(t.Block.LogsBloom),
			Miner:            str(t.Block.Miner),
			MixHash:          str(t.Block.MixHash),
			Size:             str(t.Block.Size),
			StateRoot:        str(t.Block.StateRoot),
			ReceiptsRoot:     str(t.Block.ReceiptsRoot),
			TotalDifficulty:  str(t.Block.TotalDifficulty),
			Sha3Uncles:       str(t.Block.Sha3Uncles),
			Uncles:           strArr(t.Block.Uncles),
			TransactionsRoot: str(t.Block.TransactionsRoot),
			Transactions:     txs,
		},
		Timestamps: t.Timestamps.ToMessage(),
	}, nil
}

type TransactionEvent struct {
	BlockEvt    *BlockEvent
	Transaction *Transaction
	Receipt     *TransactionReceipt
	Timestamps  *TrackingTimestamps
}

func safeAddStrValueToMap(addresses map[string]bool, addr string) {
	if addr != "" {
		addresses[strings.ToLower(addr)] = true
	}
}

func safeAddStrToMap(addresses map[string]bool, addr *string) {
	if addr != nil {
		safeAddStrValueToMap(addresses, *addr)
	}
}

// ToMessage converts the TransactionEvent to the protocol.TransactionEvent message
func (t *TransactionEvent) ToMessage() (*protocol.TransactionEvent, error) {
	evtType := protocol.TransactionEvent_BLOCK

	addresses := make(map[string]bool)

	um := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}

	// convert trace domain model to proto (filter traces)
	var traces []*protocol.TransactionEvent_Trace
	for _, trace := range t.BlockEvt.Traces {
		if trace.TransactionHash != nil && *trace.TransactionHash == t.Transaction.Hash {
			safeAddStrToMap(addresses, trace.Action.Address)
			safeAddStrToMap(addresses, trace.Action.RefundAddress)
			safeAddStrToMap(addresses, trace.Action.To)
			safeAddStrToMap(addresses, trace.Action.From)

			var pTrace protocol.TransactionEvent_Trace
			traceJson, err := json.Marshal(trace)
			if err != nil {
				return nil, err
			}
			if err := um.Unmarshal(bytes.NewReader(traceJson), &pTrace); err != nil {
				log.Errorf("cannot unmarshal traceJson: %s", err.Error())
				log.Errorf("JSON: %s", string(traceJson))
				return nil, err
			}
			// lowercase addresses
			if pTrace.Action != nil {
				pTrace.Action.To = strings.ToLower(pTrace.Action.To)
				pTrace.Action.From = strings.ToLower(pTrace.Action.From)
				pTrace.Action.RefundAddress = strings.ToLower(pTrace.Action.RefundAddress)
				pTrace.Action.Address = strings.ToLower(pTrace.Action.Address)
			}

			traces = append(traces, &pTrace)
		}
	}

	txAddresses := make(map[string]bool)

	// convert tx domain model to proto
	var tx protocol.TransactionEvent_EthTransaction
	if t.Transaction != nil {
		safeAddStrToMap(addresses, t.Transaction.To)
		safeAddStrToMap(addresses, &t.Transaction.From)
		safeAddStrToMap(txAddresses, t.Transaction.To)
		safeAddStrToMap(txAddresses, &t.Transaction.From)

		txJson, err := json.Marshal(t.Transaction)
		if err != nil {
			return nil, err
		}
		if err := um.Unmarshal(bytes.NewReader(txJson), &tx); err != nil {
			log.Errorf("cannot unmarshal txJson: %s", err.Error())
			log.Errorf("JSON: %s", string(txJson))
			return nil, err
		}

		// lowercase to/from
		tx.To = strings.ToLower(tx.To)
		tx.From = strings.ToLower(tx.From)
	}

	var logs []*protocol.TransactionEvent_Log
	logJson, err := json.Marshal(t.BlockEvt.Logs)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(logJson, &logs); err != nil {
		return nil, err
	}

	var txLogs []*protocol.TransactionEvent_Log
	for _, l := range logs {
		if l.TransactionHash == t.Transaction.Hash {
			txLogs = append(txLogs, l)
			l.Address = strings.ToLower(l.Address)
			safeAddStrValueToMap(addresses, l.Address)
			safeAddStrValueToMap(txAddresses, l.Address)

			// add addresses from topics
			for _, topic := range l.Topics {
				if strings.HasPrefix(topic, "0x000000000000000000000000") {
					safeAddStrValueToMap(addresses, common.HexToAddress(topic).Hex())
					safeAddStrValueToMap(txAddresses, common.HexToAddress(topic).Hex())
				}
			}
		}
	}

	contractAddress := ""
	isDeploy := t.Transaction.To == nil
	if isDeploy {
		createdAddr := crypto.CreateAddress(common.HexToAddress(t.Transaction.From), uint64(utils.HexToInt64(t.Transaction.Nonce)))
		contractAddress = strings.ToLower(createdAddr.Hex())
		safeAddStrValueToMap(addresses, contractAddress)
		safeAddStrValueToMap(txAddresses, contractAddress)
	}

	// for backwards compatibility
	receipt := protocol.TransactionEvent_EthReceipt{
		Root:              "",
		Status:            "0x1",
		CumulativeGasUsed: "",
		LogsBloom:         "",
		Logs:              txLogs,
		TransactionHash:   t.Transaction.Hash,
		ContractAddress:   contractAddress,
		GasUsed:           t.Transaction.Gas,
		BlockHash:         t.BlockEvt.Block.Hash,
		BlockNumber:       t.BlockEvt.Block.Number,
		TransactionIndex:  t.Transaction.TransactionIndex,
	}

	nw := &protocol.TransactionEvent_Network{}
	if t.BlockEvt.ChainID != nil {
		nw.ChainId = utils.BigIntToHex(t.BlockEvt.ChainID)
	}

	return &protocol.TransactionEvent{
		Type:                 evtType,
		Transaction:          &tx,
		Network:              nw,
		Traces:               traces,
		Addresses:            addresses,
		TxAddresses:          txAddresses,
		Logs:                 txLogs,
		Receipt:              &receipt,
		IsContractDeployment: isDeploy,
		ContractAddress:      contractAddress,
		Block: &protocol.TransactionEvent_EthBlock{
			BlockHash:      t.BlockEvt.Block.Hash,
			BlockNumber:    t.BlockEvt.Block.Number,
			BlockTimestamp: t.BlockEvt.Block.Timestamp,
		},
		Timestamps: t.Timestamps.ToMessage(),
	}, nil
}
