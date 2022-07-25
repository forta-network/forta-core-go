package inspect

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/hashicorp/go-multierror"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

// validationError contains validation error data.
type validationError struct {
	ErrorCode    string
	ErrorMessage string
}

// compile time check: it should implement the interface
var _ error = &validationError{}

func (err *validationError) Error() string {
	return fmt.Sprintf("%s (code: %s)", err.ErrorMessage, err.ErrorCode)
}

func (err *validationError) Code() string {
	return err.ErrorCode
}

func (err *validationError) Message() string {
	return err.ErrorMessage
}

// Validation errors
var (
	// 1xxx: validator messed up at initialization
	ErrReferenceScanAPI = &validationError{
		ErrorCode:    "1001",
		ErrorMessage: "reference scan api is failing",
	}
	ErrReferenceTraceAPI = &validationError{
		ErrorCode:    "1002",
		ErrorMessage: "reference trace api is failing",
	}
	ErrReferenceProxyAPI = &validationError{
		ErrorCode:    "1003",
		ErrorMessage: "reference proxy api is failing",
	}

	// 2xxx: validator messed up (bad sources etc.)
	// 2x0x: validator could not fetch reference data
	// 2x1x: validator retrieved bad/mismatching data
	ErrReferenceScanAPIBlock = &validationError{
		ErrorCode:    "2001",
		ErrorMessage: "reference scan api could not retrieve the configured block",
	}
	ErrReferenceTraceAPIBlock = &validationError{
		ErrorCode:    "2002",
		ErrorMessage: "reference trace api could not retrieve the configured block",
	}
	ErrReferenceTraceAPITraceBlock = &validationError{
		ErrorCode:    "2003",
		ErrorMessage: "reference trace api could not retrieve the trace data for the configured block",
	}
	ErrReferenceProxyAPIBlock = &validationError{
		ErrorCode:    "2004",
		ErrorMessage: "reference proxy api could not retrieve the configured block",
	}
	ErrReferenceBlockMismatch = &validationError{
		ErrorCode:    "2011",
		ErrorMessage: "reference scan, trace and proxy apis return different blocks",
	}

	// 3xxx: node messed up: inspection results are bad
	// 3x0x: inspection results have bad/mismatching data
	// 3x1x: inspection results do not match with the reference
	ErrResultBlockMismatch = &validationError{
		ErrorCode:    "3001",
		ErrorMessage: "scan, trace and proxy inspection detected different blocks",
	}
	ErrResultScanAPIBlockMismatch = &validationError{
		ErrorCode:    "3011",
		ErrorMessage: "scan inspection detected a different block than reference",
	}
	ErrResultTraceAPIBlockMismatch = &validationError{
		ErrorCode:    "3012",
		ErrorMessage: "trace inspection detected a different block than reference",
	}
	ErrResultTraceAPITraceBlockMismatch = &validationError{
		ErrorCode:    "3013",
		ErrorMessage: "trace inspection detected different trace data than reference",
	}
	ErrResultProxyAPIBlockMismatch = &validationError{
		ErrorCode:    "3014",
		ErrorMessage: "proxy inspection detected a different block than reference",
	}
)

const (
	cacheExpiryDuration = time.Minute * 10
)

// InspectionValidator validates inspection results.
type InspectionValidator struct {
	inspectionCfg  *InspectionConfig
	scanRpcClient  *rpc.Client
	traceRpcClient *rpc.Client
	traceClient    ethereum.Client
	proxyRpcClient *rpc.Client

	cache *cache.Cache
}

// NewValidator creates a new inspection validator.
func NewValidator(ctx context.Context, inspectionCfg InspectionConfig) (*InspectionValidator, error) {
	var (
		validator InspectionValidator
		err       error
	)
	validator.scanRpcClient, err = rpc.DialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		log.WithError(err).Error("failed to dial scan api")
		return nil, ErrReferenceScanAPI
	}
	if inspectionCfg.CheckTrace {
		validator.traceRpcClient, err = rpc.DialContext(ctx, inspectionCfg.TraceAPIURL)
		if err != nil {
			log.WithError(err).Error("failed to dial trace api")
			return nil, ErrReferenceTraceAPI
		}
		validator.traceClient, err = ethereum.NewStreamEthClient(ctx, "trace", inspectionCfg.TraceAPIURL)
		if err != nil {
			log.WithError(err).Error("failed to dial trace api (NewStreamEthClient)")
			return nil, ErrReferenceTraceAPI
		}
	}
	validator.proxyRpcClient, err = rpc.DialContext(ctx, inspectionCfg.ProxyAPIURL)
	if err != nil {
		log.WithError(err).Error("failed to dial proxy api")
		return nil, ErrReferenceProxyAPI
	}
	validator.inspectionCfg = &inspectionCfg

	validator.cache = cache.New(cacheExpiryDuration, cache.DefaultExpiration)

	return &validator, nil
}

type blockValidationData struct {
	scan          string
	proxy         string
	trace         string
	traceResponse string
}

// Validate validates the inspection result. This is intended for using as a first check.
// If this validation fails, nothing else should matter because it is the first and foremost requirement
// to run a node with reliable chain data sources.
func (v *InspectionValidator) Validate(ctx context.Context, results *InspectionResults) (resultErr error) {
	var referenceBlock blockValidationData
	blockNumberStr := strconv.FormatUint(v.inspectionCfg.BlockNumber, 10)

	if item, ok := v.cache.Get(blockNumberStr); ok {
		referenceBlock = item.(blockValidationData)
	} else {
		var err error
		referenceBlock.scan, err = getBlockResponseHash(ctx, v.scanRpcClient, v.inspectionCfg.BlockNumber)
		if err != nil {
			log.WithError(err).Error("failed to get scan api block response hash")
			resultErr = multierror.Append(resultErr, ErrReferenceScanAPI)
		}

		if v.inspectionCfg.CheckTrace {
			referenceBlock.trace, err = getBlockResponseHash(ctx, v.traceRpcClient, v.inspectionCfg.BlockNumber)
			if err != nil {
				log.WithError(err).Error("failed to get trace api block response hash")
				resultErr = multierror.Append(resultErr, ErrReferenceTraceAPI)
			}
			referenceBlock.traceResponse, err = getTraceResponseHash(ctx, v.traceClient, v.inspectionCfg.BlockNumber)
			if err != nil {
				log.WithError(err).Error("failed to get trace api trace block response hash")
				resultErr = multierror.Append(resultErr, ErrReferenceTraceAPI)
			}
		}

		referenceBlock.proxy, err = getBlockResponseHash(ctx, v.proxyRpcClient, v.inspectionCfg.BlockNumber)
		if err != nil {
			log.WithError(err).Error("failed to get proxy api block response hash")
			resultErr = multierror.Append(resultErr, ErrReferenceProxyAPI)
		}

		v.cache.Set(blockNumberStr, referenceBlock, cacheExpiryDuration)
	}

	// return with the first batch of errors at this step
	if resultErr != nil {
		return
	}

	// check if this validator is messing up
	if (referenceBlock.scan != referenceBlock.trace && v.inspectionCfg.CheckTrace) ||
		referenceBlock.scan != referenceBlock.proxy {
		resultErr = multierror.Append(resultErr, ErrReferenceBlockMismatch)
	}

	// check if the inspected node is messing up
	traceApiBlockHashIsDifferent := results.Metadata[MetadataScanAPIBlockByNumberHash] != results.Metadata[MetadataTraceAPIBlockByNumberHash] &&
		v.inspectionCfg.CheckTrace

	proxyApiBlockHashIsDifferent := results.Metadata[MetadataScanAPIBlockByNumberHash] != results.Metadata[MetadataProxyAPIBlockByNumberHash]

	if traceApiBlockHashIsDifferent || proxyApiBlockHashIsDifferent {
		resultErr = multierror.Append(resultErr, ErrResultBlockMismatch)
	}
	if results.Metadata[MetadataScanAPIBlockByNumberHash] != referenceBlock.scan {
		resultErr = multierror.Append(resultErr, ErrResultScanAPIBlockMismatch)
	}
	if v.inspectionCfg.CheckTrace {
		if results.Metadata[MetadataTraceAPIBlockByNumberHash] != referenceBlock.trace {
			resultErr = multierror.Append(resultErr, ErrResultTraceAPIBlockMismatch)
		}
		if results.Metadata[MetadataTraceAPITraceBlockHash] != referenceBlock.traceResponse {
			resultErr = multierror.Append(resultErr, ErrResultTraceAPITraceBlockMismatch)
		}
	}
	if results.Metadata[MetadataProxyAPIBlockByNumberHash] != referenceBlock.proxy {
		resultErr = multierror.Append(resultErr, ErrResultProxyAPIBlockMismatch)
	}

	return
}
