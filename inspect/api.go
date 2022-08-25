package inspect

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/domain"
)

const (
	blockByNumber = "eth_getBlockByNumber"
	latestBlock   = "latest"
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

// GetNetworkID gets the network ID from net_version.
func GetNetworkID(ctx context.Context, rpcClient *rpc.Client) (*big.Int, error) {
	var resultStr string
	err := rpcClient.CallContext(ctx, &resultStr, "net_version")
	if err == nil {
		return decodeChainID(resultStr)
	}

	var result big.Int
	err = rpcClient.CallContext(ctx, &result, "net_version")
	if err == nil {
		return &result, nil
	}

	return nil, fmt.Errorf("net_version failed: %v", err)
}

// GetChainID gets the chain ID from eth_chainId.
func GetChainID(ctx context.Context, rpcClient *rpc.Client) (*big.Int, error) {
	var result string
	err := rpcClient.CallContext(ctx, &result, "eth_chainId")
	if err != nil {
		return nil, fmt.Errorf("eth_chainId failed: %v", err)
	}
	return decodeChainID(result)
}

// GetChainOrNetworkID gets the chain ID from either of eth_chainId or net_version.
func GetChainOrNetworkID(ctx context.Context, rpcClient *rpc.Client) (*big.Int, error) {
	num, err1 := GetChainID(ctx, rpcClient)
	if err1 == nil {
		return num, nil
	}
	num, err2 := GetNetworkID(ctx, rpcClient)
	if err2 == nil {
		return num, nil
	}
	return nil, fmt.Errorf("neither eth_chainId nor net_version worked: %v, %v", err1, err2)
}

func decodeChainID(numStr string) (*big.Int, error) {
	num, err := hexutil.DecodeBig(numStr)
	if err == nil {
		return num, nil
	}
	num, ok := big.NewInt(0).SetString(numStr, 10)
	if ok {
		return num, nil
	}
	return nil, fmt.Errorf("could not decode chain ID '%s' - it is neither hex nor base10", numStr)
}

func SupportsETH2(ctx context.Context, rpcClient *rpc.Client) bool {
	chainID, err := GetChainOrNetworkID(ctx, rpcClient)
	if err != nil {
		return false
	}
	if !(chainID.Uint64() == 1 || chainID.Uint64() == 5) { // eth mainnet and goerli
		return false
	}

	var block domain.Block
	if err := getRpcResponse(ctx, rpcClient, &block, blockByNumber, latestBlock, true); err != nil {
		return false
	}
	if block.Difficulty == nil {
		return false
	}
	if block.Nonce == nil {
		return false
	}
	difficulty, err := hexutil.DecodeBig(*block.Difficulty)
	if err != nil {
		return false
	}
	var nonce types.BlockNonce
	if err := (&nonce).UnmarshalText([]byte(*block.Nonce)); err != nil {
		return false
	}
	if err != nil {
		return false
	}

	if difficulty.Sign() == 0 && nonce.Uint64() == 0 {
		return true
	}
	return false
}
