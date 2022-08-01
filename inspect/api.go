package inspect

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/domain"
)

// GetBlockResponseHash computes a hash by using some data from the API response.
func GetBlockResponseHash(ctx context.Context, rpcClient *rpc.Client, blockNumber uint64) (string, error) {
	var block domain.Block
	if err := getRpcResponse(ctx, rpcClient, &block, "eth_getBlockByNumber", hexutil.EncodeUint64(blockNumber), true); err != nil {
		return "", err
	}
	hashConcat := block.Hash
	for _, tx := range block.Transactions {
		hashConcat += tx.Hash
	}
	return hashOf(hashConcat), nil
}

// GetTraceResponseHash computes a hash by using some data from the API response.
func GetTraceResponseHash(ctx context.Context, rpcClient *rpc.Client, blockNumber uint64) (string, error) {
	var traces []*domain.Trace
	if err := getRpcResponse(ctx, rpcClient, &traces, "trace_block", hexutil.EncodeUint64(blockNumber)); err != nil {
		return "", err
	}
	var hashConcat string
	for _, trace := range traces {
		if trace.TransactionHash != nil {
			hashConcat += *trace.TransactionHash
		}
	}
	return hashOf(hashConcat), nil
}

func hashOf(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func getRpcResponse(ctx context.Context, rpcClient *rpc.Client, respData interface{}, method string, args ...interface{}) error {
	return rpcClient.CallContext(ctx, &respData, method, args...)
}
