package domain

import (
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	SourceAlert time.Time
	Feed        time.Time
	BotRequest  time.Time
	BotResponse time.Time
}

func (tt *TrackingTimestamps) ToMessage() *protocol.TrackingTimestamps {
	return &protocol.TrackingTimestamps{
		Block:       tt.Block.Format(TimeTrackingTimestampFormat),
		Feed:        tt.Feed.Format(TimeTrackingTimestampFormat),
		SourceAlert: tt.SourceAlert.Format(TimeTrackingTimestampFormat),
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
		SourceAlert: TimeFromString(tt.SourceAlert),
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

func (t *BlockEvent) ToBlockData() *protocol.BlockData {
	combinedBlockEvent := &protocol.BlockData{
		ChainID: t.ChainID.Uint64(),
		Block: &protocol.BlockWithTransactions{
			BaseFeePerGas:         str(t.Block.BaseFeePerGas),
			Difficulty:            str(t.Block.Difficulty),
			ExtraData:             str(t.Block.ExtraData),
			GasLimit:              str(t.Block.GasLimit),
			GasUsed:               str(t.Block.GasUsed),
			Hash:                  t.Block.Hash,
			LogsBloom:             str(t.Block.LogsBloom),
			Miner:                 str(t.Block.Miner),
			MixHash:               str(t.Block.MixHash),
			Nonce:                 str(t.Block.Nonce),
			Number:                t.Block.Number,
			ParentHash:            t.Block.ParentHash,
			ReceiptsRoot:          str(t.Block.ReceiptsRoot),
			Size:                  str(t.Block.Size),
			StateRoot:             str(t.Block.StateRoot),
			Timestamp:             t.Block.Timestamp,
			TotalDifficulty:       str(t.Block.TotalDifficulty),
			Transactions:          make([]*protocol.Transaction, 0, len(t.Block.Transactions)),
			Sha3Uncles:            str(t.Block.Sha3Uncles),
			BlobGasUsed:           str(t.Block.BlobGasUsed),
			ParentBeaconBlockRoot: str(t.Block.ParentBeaconBlockRoot),
			ExcessBlobGas:         str(t.Block.ExcessBlobGas),
			TransactionsRoot:      str(t.Block.TransactionsRoot),
			WithdrawalsRoot:       str(t.Block.WithdrawalsRoot),
			Uncles:                make([]string, 0, len(t.Block.Uncles)),
			Withdrawals:           make([]*protocol.Withdrawal, 0, len(t.Block.Withdrawals)),
		},
		Logs:   make([]*protocol.LogEntry, 0, len(t.Logs)),
		Traces: make([]*protocol.Trace, 0, len(t.Traces)),
	}

	for _, uncle := range t.Block.Uncles {
		combinedBlockEvent.Block.Uncles = append(combinedBlockEvent.Block.Uncles, str(uncle))
	}

	for _, withdrawal := range t.Block.Withdrawals {
		combinedBlockEvent.Block.Withdrawals = append(combinedBlockEvent.Block.Withdrawals, &protocol.Withdrawal{
			Index:          withdrawal.Index,
			ValidatorIndex: withdrawal.ValidatorIndex,
			Address:        withdrawal.Address,
			Amount:         withdrawal.Amount,
		})
	}

	for i, tx := range t.Block.Transactions {
		combinedBlockEvent.Block.Transactions = append(combinedBlockEvent.Block.Transactions, &protocol.Transaction{
			From:                 tx.From,
			Gas:                  tx.Gas,
			GasPrice:             tx.GasPrice,
			Hash:                 tx.Hash,
			Input:                str(tx.Input),
			Nonce:                tx.Nonce,
			To:                   str(tx.To),
			AccessList:           make([]*protocol.AccessList, 0, len(tx.AccessList)),
			TransactionIndex:     tx.TransactionIndex,
			Value:                str(tx.Value),
			Type:                 tx.Type,
			V:                    tx.V,
			R:                    tx.R,
			S:                    tx.S,
			MaxFeePerGas:         str(tx.MaxFeePerGas),
			MaxPriorityFeePerGas: str(tx.MaxPriorityFeePerGas),
			YParity:              str(tx.YParity),
		})

		for _, accessList := range tx.AccessList {
			combinedBlockEvent.Block.Transactions[i].AccessList = append(combinedBlockEvent.Block.Transactions[i].AccessList, &protocol.AccessList{
				Address:     accessList.Address,
				StorageKeys: accessList.StorageKeys,
			})
		}
	}

	for _, log := range t.Logs {
		combinedBlockEvent.Logs = append(combinedBlockEvent.Logs, &protocol.LogEntry{
			Address:          str(log.Address),
			BlockHash:        str(log.BlockHash),
			BlockNumber:      str(log.BlockNumber),
			Data:             str(log.Data),
			LogIndex:         str(log.LogIndex),
			Removed:          safeBool(log.Removed),
			Topics:           strArr(log.Topics),
			TransactionHash:  str(log.TransactionHash),
			TransactionIndex: str(log.TransactionIndex),
		})
	}

	for _, trace := range t.Traces {
		t := &protocol.Trace{
			Action: &protocol.TraceAction{
				CallType:      str(trace.Action.CallType),
				To:            str(trace.Action.To),
				Input:         str(trace.Action.Input),
				From:          str(trace.Action.From),
				Gas:           str(trace.Action.Gas),
				Value:         str(trace.Action.Value),
				Init:          str(trace.Action.Init),
				Address:       str(trace.Action.Address),
				Balance:       str(trace.Action.Balance),
				RefundAddress: str(trace.Action.RefundAddress),
			},
			Subtraces:           int32(trace.Subtraces),
			TraceAddress:        intArr(trace.TraceAddress),
			TransactionHash:     str(trace.TransactionHash),
			TransactionPosition: int64(safeInt(trace.TransactionPosition)),
			Type:                trace.Type,
			Error:               str(trace.Error),
		}

		if trace.Result != nil {
			t.Result = &protocol.TraceResult{
				Output:  str(trace.Result.Output),
				GasUsed: str(trace.Result.GasUsed),
				Address: str(trace.Result.Address),
				Code:    str(trace.Result.Code),
			}
		}
		combinedBlockEvent.Traces = append(combinedBlockEvent.Traces, t)
	}

	return combinedBlockEvent
}

func safeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func intArr(vals []int) []int64 {
	result := make([]int64, 0, len(vals))
	for _, v := range vals {
		result = append(result, int64(v))
	}
	return result
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
			BaseFeePerGas:    str(t.Block.BaseFeePerGas),
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

func safeBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// ToMessage converts the TransactionEvent to the protocol.TransactionEvent message
func (t *TransactionEvent) ToMessage() (*protocol.TransactionEvent, error) {
	evtType := protocol.TransactionEvent_BLOCK

	addresses := make(map[string]bool)

	// convert trace domain model to proto (filter traces)
	var traces []*protocol.TransactionEvent_Trace
	for _, trace := range t.BlockEvt.Traces {
		if trace.TransactionHash != nil && *trace.TransactionHash == t.Transaction.Hash {
			safeAddStrToMap(addresses, trace.Action.Address)
			safeAddStrToMap(addresses, trace.Action.RefundAddress)
			safeAddStrToMap(addresses, trace.Action.To)
			safeAddStrToMap(addresses, trace.Action.From)

			pTrace := trace.ToProto()

			// lowercase addresses
			if pTrace.Action != nil {
				pTrace.Action.To = strings.ToLower(pTrace.Action.To)
				pTrace.Action.From = strings.ToLower(pTrace.Action.From)
				pTrace.Action.RefundAddress = strings.ToLower(pTrace.Action.RefundAddress)
				pTrace.Action.Address = strings.ToLower(pTrace.Action.Address)
			}

			traces = append(traces, pTrace)
		}
	}

	txAddresses := make(map[string]bool)

	// convert tx domain model to proto
	var tx *protocol.TransactionEvent_EthTransaction
	if t.Transaction != nil {
		safeAddStrToMap(addresses, t.Transaction.To)
		safeAddStrToMap(addresses, &t.Transaction.From)
		safeAddStrToMap(txAddresses, t.Transaction.To)
		safeAddStrToMap(txAddresses, &t.Transaction.From)

		tx = t.Transaction.ToProto()

		// lowercase to/from
		tx.To = strings.ToLower(tx.To)
		tx.From = strings.ToLower(tx.From)
	}

	var txLogs []*protocol.TransactionEvent_Log
	for _, l := range t.BlockEvt.Logs {
		if l.TransactionHash != nil && *l.TransactionHash == t.Transaction.Hash {
			txLog := &protocol.TransactionEvent_Log{
				Address:          str(l.Address),
				Topics:           strArr(l.Topics),
				Data:             str(l.Data),
				BlockNumber:      str(l.BlockNumber),
				TransactionHash:  str(l.TransactionHash),
				TransactionIndex: str(l.TransactionIndex),
				BlockHash:        str(l.BlockHash),
				LogIndex:         str(l.LogIndex),
				Removed:          safeBool(l.Removed),
			}

			txLogs = append(txLogs, txLog)
			txLog.Address = strings.ToLower(txLog.Address)
			safeAddStrValueToMap(addresses, txLog.Address)
			safeAddStrValueToMap(txAddresses, txLog.Address)

			// add addresses from topics
			for _, topic := range txLog.Topics {
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
		Transaction:          tx,
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
			BaseFeePerGas:  str(t.BlockEvt.Block.BaseFeePerGas),
		},
		Timestamps: t.Timestamps.ToMessage(),
	}, nil
}

type AlertEvent struct {
	EventType  EventType
	Event      *protocol.AlertEvent
	Logs       []LogEntry
	Timestamps *TrackingTimestamps
	Subscriber Subscriber
}

// ToMessage converts the AlertEvent to the protocol.TransactionEvent message
func (t *AlertEvent) ToMessage() (*protocol.AlertEvent, error) {
	return &protocol.AlertEvent{
		Alert:      t.Event.Alert,
		Timestamps: t.Timestamps.ToMessage(),
	}, nil
}

type Subscriber struct {
	BotID    string `json:"bot_id"`
	BotOwner string `json:"bot_owner"`
	BotImage string `json:"bot_image_hash"`
}

type CombinerBotSubscription struct {
	Subscription *protocol.CombinerBotSubscription `json:"subscription"`
	Subscriber   *Subscriber                       `json:"subscriber"`
}

// Equal returns true if two subscriptions are equal and false otherwise.
func (c *CombinerBotSubscription) Equal(b *CombinerBotSubscription) bool {
	if c == nil || b == nil {
		return false
	}

	if c.Subscription.BotId != b.Subscription.BotId ||
		c.Subscription.AlertId != b.Subscription.AlertId ||
		c.Subscription.ChainId != b.Subscription.ChainId ||
		!stringSlicesEqual(c.Subscription.AlertIds, b.Subscription.AlertIds) {
		return false
	}

	// Subscriber-specific uniqueness checks. Since the protocol enforces subscription fees, subscriptions from 2 different bots or bot owners can not
	// be treated as same, because one can fail while the other succeeds.
	if c.Subscriber.BotID != b.Subscriber.BotID || c.Subscriber.BotOwner != b.Subscriber.BotOwner || c.Subscriber.BotImage != b.Subscriber.BotImage {
		return false
	}

	return true
}

// stringSlicesEqual returns true if two slices of strings are equal and false otherwise.
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Strings(aCopy)
	sort.Strings(bCopy)

	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}

	return true
}
