package domain

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
)

// Block is the intersection between parity and go-ethereum block
type Block struct {
	BaseFeePerGas         *string       `json:"baseFeePerGas"`
	BlobGasUsed           *string       `json:"blobGasUsed"`
	Difficulty            *string       `json:"difficulty"`
	ExcessBlobGas         *string       `json:"excessBlobGas"`
	ExtraData             *string       `json:"extraData"`
	GasLimit              *string       `json:"gasLimit"`
	GasUsed               *string       `json:"gasUsed"`
	Hash                  string        `json:"hash"`
	LogsBloom             *string       `json:"logsBloom"`
	Miner                 *string       `json:"miner"`
	MixHash               *string       `json:"mixHash"`
	Nonce                 *string       `json:"nonce"`
	Number                string        `json:"number"`
	ParentBeaconBlockRoot *string       `json:"parentBeaconBlockRoot"`
	ParentHash            string        `json:"parentHash"`
	ReceiptsRoot          *string       `json:"receiptsRoot"`
	Sha3Uncles            *string       `json:"sha3Uncles"`
	Size                  *string       `json:"size"`
	StateRoot             *string       `json:"stateRoot"`
	Timestamp             string        `json:"timestamp"`
	TotalDifficulty       *string       `json:"totalDifficulty"`
	Transactions          []Transaction `json:"transactions"`
	TransactionsRoot      *string       `json:"transactionsRoot"`
	Uncles                []*string     `json:"uncles"`
	Withdrawals           []Withdrawal  `json:"withdrawals"`
	WithdrawalsRoot       *string       `json:"withdrawalsRoot"`
}

type Withdrawal struct {
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
}

func (b *Block) Age() (*time.Duration, error) {
	ts, err := b.GetTimestamp()
	if err != nil {
		return nil, err
	}
	age := time.Since(*ts)
	return &age, nil
}

func (b *Block) GetTimestamp() (*time.Time, error) {
	ts, err := utils.HexToBigInt(b.Timestamp)
	if err != nil {
		return nil, err
	}
	blockTsMs := ts.Mul(ts, big.NewInt(1000))
	result := time.Unix(0, int64(blockTsMs.Uint64())*int64(time.Millisecond))
	return &result, nil
}

func BlockFromBlockData(bd *protocol.BlockData) *Block {
	block := &Block{
		Difficulty:            &bd.Block.Difficulty,
		ExtraData:             &bd.Block.ExtraData,
		GasLimit:              &bd.Block.GasLimit,
		GasUsed:               &bd.Block.GasUsed,
		Hash:                  bd.Block.Hash,
		LogsBloom:             &bd.Block.LogsBloom,
		Miner:                 &bd.Block.Miner,
		MixHash:               &bd.Block.MixHash,
		Nonce:                 &bd.Block.Nonce,
		Number:                bd.Block.Number,
		ParentHash:            bd.Block.ParentHash,
		ReceiptsRoot:          &bd.Block.ReceiptsRoot,
		ParentBeaconBlockRoot: &bd.Block.ParentBeaconBlockRoot,
		Sha3Uncles:            &bd.Block.Sha3Uncles,
		Size:                  &bd.Block.Size,
		StateRoot:             &bd.Block.StateRoot,
		Timestamp:             bd.Block.Timestamp,
		TotalDifficulty:       &bd.Block.TotalDifficulty,
		TransactionsRoot:      &bd.Block.TransactionsRoot,
		BlobGasUsed:           &bd.Block.BlobGasUsed,
		BaseFeePerGas:         &bd.Block.BaseFeePerGas,
		WithdrawalsRoot:       &bd.Block.WithdrawalsRoot,
		ExcessBlobGas:         &bd.Block.ExcessBlobGas,
	}

	for _, uncle := range bd.Block.Uncles {
		block.Uncles = append(block.Uncles, &uncle)
	}

	for _, withdrawal := range bd.Block.Withdrawals {
		block.Withdrawals = append(block.Withdrawals, Withdrawal{
			Index:          withdrawal.Index,
			ValidatorIndex: withdrawal.ValidatorIndex,
			Address:        withdrawal.Address,
			Amount:         withdrawal.Amount,
		})
	}

	chainID := fmt.Sprintf("0x%x", bd.ChainID)
	for i, tx := range bd.Block.Transactions {
		block.Transactions = append(block.Transactions, Transaction{
			BlockHash:            bd.Block.Hash,
			BlockNumber:          bd.Block.Number,
			From:                 tx.From,
			Gas:                  tx.Gas,
			GasPrice:             tx.GasPrice,
			Hash:                 tx.Hash,
			Input:                &tx.Input,
			Nonce:                tx.Nonce,
			To:                   &tx.To,
			TransactionIndex:     tx.TransactionIndex,
			Value:                &tx.Value,
			Type:                 tx.Type,
			ChainID:              &chainID,
			V:                    tx.V,
			R:                    tx.R,
			S:                    tx.S,
			MaxFeePerGas:         &tx.MaxFeePerGas,
			MaxPriorityFeePerGas: &tx.MaxPriorityFeePerGas,
			YParity:              &tx.YParity,
		})

		for _, access := range tx.AccessList {
			block.Transactions[i].AccessList = append(block.Transactions[i].AccessList, AccessList{
				Address:     access.Address,
				StorageKeys: access.StorageKeys,
			})
		}
	}

	return block
}

// Transaction is the intersection between parity and go-ethereum transactions
type Transaction struct {
	BlockHash            string       `json:"blockHash"`
	BlockNumber          string       `json:"blockNumber"`
	From                 string       `json:"from"`
	Gas                  string       `json:"gas"`
	GasPrice             string       `json:"gasPrice"`
	MaxPriorityFeePerGas *string      `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         *string      `json:"maxFeePerGas,omitempty"`
	Hash                 string       `json:"hash"`
	Input                *string      `json:"input"`
	Nonce                string       `json:"nonce"`
	To                   *string      `json:"to"`
	TransactionIndex     string       `json:"transactionIndex"`
	Value                *string      `json:"value"`
	Type                 string       `json:"type"`
	AccessList           []AccessList `json:"accessList,omitempty"`
	ChainID              *string      `json:"chainId"`
	V                    string       `json:"v"`
	R                    string       `json:"r"`
	S                    string       `json:"s"`
	YParity              *string      `json:"yParity"`
}

type AccessList struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

func (t *Transaction) ToProto() *protocol.TransactionEvent_EthTransaction {
	return &protocol.TransactionEvent_EthTransaction{
		Type:                 t.Type,
		Nonce:                t.Nonce,
		GasPrice:             t.GasPrice,
		Gas:                  t.Gas,
		Value:                safeValueToPointer(t.Value),
		Input:                safeValueToPointer(t.Input),
		V:                    t.V,
		R:                    t.R,
		S:                    t.S,
		To:                   safeValueToPointer(t.To),
		Hash:                 t.Hash,
		From:                 t.From,
		MaxFeePerGas:         safeValueToPointer(t.MaxFeePerGas),
		MaxPriorityFeePerGas: safeValueToPointer(t.MaxPriorityFeePerGas),
	}
}

func safeValueToPointer[T any](pointer *T) T {
	var result T
	if pointer == nil {
		return result
	}

	return *pointer
}

// LogEntry is a log item inside a receipt
type LogEntry struct {
	Address          *string   `json:"address"`
	Topics           []*string `json:"topics"`
	Data             *string   `json:"data"`
	BlockNumber      *string   `json:"blockNumber"`
	TransactionHash  *string   `json:"transactionHash"`
	TransactionIndex *string   `json:"transactionIndex"`
	BlockHash        *string   `json:"blockHash"`
	LogIndex         *string   `json:"logIndex"`
	Removed          *bool     `json:"removed"`
}

// ToTypesLog converts our type to go-ethereum type.
func (le LogEntry) ToTypesLog() (log types.Log) {
	if le.Address != nil {
		log.Address = common.HexToAddress(*le.Address)
	}
	if le.BlockHash != nil {
		log.BlockHash = common.HexToHash(*le.BlockHash)
	}
	if le.BlockNumber != nil {
		num, _ := hexutil.DecodeBig(*le.BlockNumber)
		log.BlockNumber = num.Uint64()
	}
	if le.Data != nil {
		log.Data = []byte(*le.Data)
	}
	if le.LogIndex != nil {
		num, _ := hexutil.DecodeBig(*le.LogIndex)
		log.Index = uint(num.Uint64())
	}
	if le.Removed != nil {
		log.Removed = *le.Removed
	}
	for _, topic := range le.Topics {
		log.Topics = append(log.Topics, common.HexToHash(*topic))
	}
	if le.TransactionHash != nil {
		log.TxHash = common.HexToHash(*le.TransactionHash)
	}
	if le.TransactionIndex != nil {
		num, _ := hexutil.DecodeBig(*le.TransactionIndex)
		log.TxIndex = uint(num.Uint64())
	}
	return
}

func LogsFromBlockData(bd *protocol.BlockData) []LogEntry {
	logs := make([]LogEntry, len(bd.Logs))
	for i, log := range bd.Logs {
		logs[i] = LogEntry{
			Address:          &log.Address,
			BlockHash:        &log.BlockHash,
			BlockNumber:      &log.BlockNumber,
			Data:             &log.Data,
			LogIndex:         &log.LogIndex,
			Removed:          &log.Removed,
			Topics:           make([]*string, 0, len(log.Topics)),
			TransactionHash:  &log.TransactionHash,
			TransactionIndex: &log.TransactionIndex,
		}

		for _, topic := range log.Topics {
			logs[i].Topics = append(logs[i].Topics, &topic)
		}
	}

	return logs
}

// TransactionReceipt is a result of a eth_getTransactionReceipt call
type TransactionReceipt struct {
	BlockHash         *string    `json:"blockHash"`
	BlockNumber       *string    `json:"blockNumber"`
	ContractAddress   *string    `json:"contractAddress"`
	CumulativeGasUsed *string    `json:"cumulativeGasUsed"`
	From              *string    `json:"from"`
	GasUsed           *string    `json:"gasUsed"`
	Logs              []LogEntry `json:"logs"`
	LogsBloom         *string    `json:"logsBloom"`
	Status            *string    `json:"status"`
	To                *string    `json:"to"`
	TransactionHash   *string    `json:"transactionHash"`
	TransactionIndex  *string    `json:"transactionIndex"`
}

// TraceAction is an element of a trace_block Trace response
type TraceAction struct {
	From          *string `json:"from"`
	CallType      *string `json:"callType"`
	Gas           *string `json:"gas"`
	Input         *string `json:"input"`
	To            *string `json:"to"`
	Value         *string `json:"value"`
	Init          *string `json:"init"`
	Address       *string `json:"address"`
	Balance       *string `json:"balance"`
	RefundAddress *string `json:"refundAddress"`
}

// TraceResult is a result element of a trace_block Trace response
type TraceResult struct {
	GasUsed *string `json:"gasUsed"`
	Output  *string `json:"output"`
	Address *string `json:"address"`
	Code    *string `json:"code"`
}

// Trace is a specific traced action in a transaction
type Trace struct {
	Action              TraceAction  `json:"action"`
	BlockHash           *string      `json:"blockHash"`
	BlockNumber         *int         `json:"blockNumber"`
	Error               *string      `json:"error"`
	Result              *TraceResult `json:"result"`
	Subtraces           int          `json:"subtraces"`
	TraceAddress        []int        `json:"traceAddress"`
	TransactionHash     *string      `json:"transactionHash"`
	TransactionPosition *int         `json:"transactionPosition"`
	Type                string       `json:"type"`
}

func (t Trace) ToProto() *protocol.TransactionEvent_Trace {
	traceAddress := make([]int64, len(t.TraceAddress))
	for i, address := range t.TraceAddress {
		traceAddress[i] = int64(address)
	}
	var traceResult *protocol.TransactionEvent_TraceResult
	if t.Result != nil {
		traceResult = &protocol.TransactionEvent_TraceResult{
			GasUsed: safeValueToPointer(t.Result.GasUsed),
			Address: safeValueToPointer(t.Result.Address),
			Code:    safeValueToPointer(t.Result.Code),
			Output:  safeValueToPointer(t.Result.Output),
		}
	}
	return &protocol.TransactionEvent_Trace{
		Action: &protocol.TransactionEvent_TraceAction{
			CallType:      safeValueToPointer(t.Action.CallType),
			To:            safeValueToPointer(t.Action.To),
			Input:         safeValueToPointer(t.Action.Input),
			From:          safeValueToPointer(t.Action.From),
			Value:         safeValueToPointer(t.Action.Value),
			Init:          safeValueToPointer(t.Action.Init),
			Address:       safeValueToPointer(t.Action.Address),
			Balance:       safeValueToPointer(t.Action.Balance),
			RefundAddress: safeValueToPointer(t.Action.RefundAddress),
		},
		BlockHash:           safeValueToPointer(t.BlockHash),
		BlockNumber:         int64(safeValueToPointer(t.BlockNumber)),
		Result:              traceResult,
		Subtraces:           int64(t.Subtraces),
		TraceAddress:        traceAddress,
		TransactionHash:     safeValueToPointer(t.TransactionHash),
		TransactionPosition: int64(safeValueToPointer(t.TransactionPosition)),
		Type:                t.Type,
		Error:               safeValueToPointer(t.Error),
	}
}

func TracesFromBlockData(bd *protocol.BlockData) ([]Trace, error) {
	traces := make([]Trace, len(bd.Traces))
	blockNumber, err := strconv.ParseInt(strings.Replace(bd.Block.Number, "0x", "", -1), 16, 64)
	if err != nil {
		return nil, err
	}

	intBlockNumber := int(blockNumber)
	for i, trace := range bd.Traces {
		transactionPosition := int(trace.TransactionPosition)
		traceAddresses := make([]int, len(trace.TraceAddress))
		for j, address := range trace.TraceAddress {
			traceAddresses[j] = int(address)
		}

		traces[i] = Trace{
			BlockNumber:         &intBlockNumber,
			Subtraces:           int(trace.Subtraces),
			TraceAddress:        traceAddresses,
			TransactionHash:     &trace.TransactionHash,
			TransactionPosition: &transactionPosition,
			Type:                trace.Type,
			Error:               &trace.Error,
		}

		if bd.Block != nil {
			traces[i].BlockHash = &bd.Block.Hash
		}

		if trace.Action != nil {
			traces[i].Action = TraceAction{
				CallType:      &trace.Action.CallType,
				To:            &trace.Action.To,
				Input:         &trace.Action.Input,
				From:          &trace.Action.From,
				Gas:           &trace.Action.Gas,
				Value:         &trace.Action.Value,
				Init:          &trace.Action.Init,
				Address:       &trace.Action.Address,
				Balance:       &trace.Action.Balance,
				RefundAddress: &trace.Action.RefundAddress,
			}
		}

		if trace.Result != nil {
			traces[i].Result = &TraceResult{
				Output:  &trace.Result.Output,
				GasUsed: &trace.Result.GasUsed,
				Address: &trace.Result.Address,
				Code:    &trace.Result.Code,
			}
		}
	}

	return traces, nil
}

// HeaderCh provides new block headers.
type HeaderCh <-chan *types.Header

// ClientSubscription abstracts away the subscription implementation.
type ClientSubscription interface {
	Err() <-chan error
	Unsubscribe()
}
