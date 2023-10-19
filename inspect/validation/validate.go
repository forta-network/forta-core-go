package validation

import (
	"context"
	"strconv"
	"time"

	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/inspect"
	"github.com/hashicorp/go-multierror"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

const (
	cacheExpiryDuration = time.Minute * 10
)

// InspectionValidator validates inspection results.
type InspectionValidator struct {
	inspectionCfg  *inspect.InspectionConfig
	scanRpcClient  ethereum.RPCClient
	traceRpcClient ethereum.RPCClient
	proxyRpcClient ethereum.RPCClient

	cache *cache.Cache
}

// NewValidator creates a new inspection validator.
func NewValidator(ctx context.Context, inspectionCfg inspect.InspectionConfig) (*InspectionValidator, error) {
	var (
		validator InspectionValidator
		err       error
	)
	validator.scanRpcClient, err = inspect.RPCDialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		log.WithError(err).Error("failed to dial scan api")
		return nil, inspect.ErrReferenceScanAPI
	}
	if inspectionCfg.CheckTrace {
		validator.traceRpcClient, err = inspect.RPCDialContext(ctx, inspectionCfg.TraceAPIURL)
		if err != nil {
			log.WithError(err).Error("failed to dial trace api")
			return nil, inspect.ErrReferenceTraceAPI
		}
	}
	validator.proxyRpcClient, err = inspect.RPCDialContext(ctx, inspectionCfg.ProxyAPIURL)
	if err != nil {
		log.WithError(err).Error("failed to dial proxy api")
		return nil, inspect.ErrReferenceProxyAPI
	}
	validator.inspectionCfg = &inspectionCfg

	validator.cache = cache.New(cacheExpiryDuration, cache.DefaultExpiration)

	return &validator, nil
}

// HashReferences contains hash references that are used during inspection validation.
type HashReferences struct {
	ScanAPIBlockHash  string
	ProxyAPIBlockHash string
	TraceAPIBlockHash string
	TraceAPITraceHash string
}

// Validate validates the inspection result. This is intended for using as a first check.
// If this validation fails, nothing else should matter because it is the first and foremost requirement
// to run a node with reliable chain data sources.
func (v *InspectionValidator) Validate(ctx context.Context, results *inspect.InspectionResults) (validationErrs inspect.ValidationErrors, resultErr error) {
	refData, err := v.getReferenceData(ctx, results)
	if err != nil {
		validationErrs = inspect.ValidationErrorsFrom(err)
		resultErr = err
		return
	}
	return inspect.ValidateHashReferences(results.Metadata, results.Inputs, &refData)
}

func (v *InspectionValidator) getReferenceData(ctx context.Context, results *inspect.InspectionResults) (refData inspect.HashReferences, resultErr error) {
	blockNumber := results.Inputs.BlockNumber
	blockNumberStr := strconv.FormatUint(blockNumber, 10)
	if item, ok := v.cache.Get(blockNumberStr); ok {
		return item.(inspect.HashReferences), nil
	}

	var err error
	refData.ScanAPIBlockHash, err = inspect.GetBlockResponseHash(ctx, v.scanRpcClient, blockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get scan api block response hash")
		resultErr = multierror.Append(resultErr, inspect.ErrReferenceScanAPIBlock)
	}

	if v.inspectionCfg.CheckTrace {
		refData.TraceAPIBlockHash, err = inspect.GetBlockResponseHash(ctx, v.traceRpcClient, blockNumber)
		if err != nil {
			log.WithError(err).Error("failed to get trace api block response hash")
			resultErr = multierror.Append(resultErr, inspect.ErrReferenceTraceAPIBlock)
		}
		refData.TraceAPITraceHash, err = inspect.GetTraceResponseHash(ctx, v.traceRpcClient, blockNumber)
		if err != nil {
			log.WithError(err).Error("failed to get trace api trace block response hash")
			resultErr = multierror.Append(resultErr, inspect.ErrReferenceTraceAPITraceBlock)
		}
	}

	refData.ProxyAPIBlockHash, err = inspect.GetBlockResponseHash(ctx, v.proxyRpcClient, blockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get proxy api block response hash")
		resultErr = multierror.Append(resultErr, inspect.ErrReferenceProxyAPIBlock)
	}

	v.cache.Set(blockNumberStr, refData, cacheExpiryDuration)
	return
}
