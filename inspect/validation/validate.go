package validation

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/inspect"
	"github.com/hashicorp/go-multierror"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

// validationError contains validation error data.
type validationError struct {
	ErrorCode    string
	ErrorMessage string
}

// ValidationError is the inspection validation error interface.
type ValidationError interface {
	Error() string
	Code() string
	Message() string
}

// ValidationErrors is a type alias which lets us implement convenience methods.
type ValidationErrors []ValidationError

// Codes returns an array of error codes.
func (verrs ValidationErrors) Codes() (codes []string) {
	if len(verrs) == 0 {
		return
	}
	for _, verr := range verrs {
		codes = append(codes, verr.Code())
	}
	return
}

// HasCode checks if the validation errors has an error with this code.
func (verrs ValidationErrors) HasCode(code string) bool {
	for _, verr := range verrs {
		if verr.Code() == code {
			return true
		}
	}
	return false
}

// HasCodes checks if the validation errors has any of these codes.
func (verrs ValidationErrors) HasCodes(codes ...string) bool {
	for _, code := range codes {
		if verrs.HasCode(code) {
			return true
		}
	}
	return false
}

// compile time check: it should implement the interfaces
var _ error = &validationError{}
var _ ValidationError = &validationError{}

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
	ErrReferenceScanAPI ValidationError = &validationError{
		ErrorCode:    "1001",
		ErrorMessage: "reference scan api is failing",
	}
	ErrReferenceTraceAPI ValidationError = &validationError{
		ErrorCode:    "1002",
		ErrorMessage: "reference trace api is failing",
	}
	ErrReferenceProxyAPI ValidationError = &validationError{
		ErrorCode:    "1003",
		ErrorMessage: "reference proxy api is failing",
	}

	// 2xxx: validator messed up (bad sources etc.)
	// 2x0x: validator could not fetch reference data
	// 2x1x: validator retrieved bad/mismatching data
	ErrReferenceScanAPIBlock ValidationError = &validationError{
		ErrorCode:    "2001",
		ErrorMessage: "reference scan api could not retrieve the configured block",
	}
	ErrReferenceTraceAPIBlock ValidationError = &validationError{
		ErrorCode:    "2002",
		ErrorMessage: "reference trace api could not retrieve the configured block",
	}
	ErrReferenceTraceAPITraceBlock ValidationError = &validationError{
		ErrorCode:    "2003",
		ErrorMessage: "reference trace api could not retrieve the trace data for the configured block",
	}
	ErrReferenceProxyAPIBlock ValidationError = &validationError{
		ErrorCode:    "2004",
		ErrorMessage: "reference proxy api could not retrieve the configured block",
	}
	ErrReferenceBlockMismatch ValidationError = &validationError{
		ErrorCode:    "2011",
		ErrorMessage: "reference scan, trace and proxy apis return different blocks",
	}

	// 3xxx: node messed up: inspection results are bad
	// 3x0x: inspection results have bad/mismatching data
	// 3x1x: inspection results do not match with the reference
	ErrResultBlockMismatch ValidationError = &validationError{
		ErrorCode:    "3001",
		ErrorMessage: "scan, trace and proxy inspection detected different blocks",
	}
	ErrResultScanAPIBlockMismatch ValidationError = &validationError{
		ErrorCode:    "3011",
		ErrorMessage: "scan inspection detected a different block than reference",
	}
	ErrResultTraceAPIBlockMismatch ValidationError = &validationError{
		ErrorCode:    "3012",
		ErrorMessage: "trace inspection detected a different block than reference",
	}
	ErrResultTraceAPITraceBlockMismatch ValidationError = &validationError{
		ErrorCode:    "3013",
		ErrorMessage: "trace inspection detected different trace data than reference",
	}
	ErrResultProxyAPIBlockMismatch ValidationError = &validationError{
		ErrorCode:    "3014",
		ErrorMessage: "proxy inspection detected a different block than reference",
	}
)

const (
	cacheExpiryDuration = time.Minute * 10
)

// InspectionValidator validates inspection results.
type InspectionValidator struct {
	inspectionCfg  *inspect.InspectionConfig
	scanRpcClient  *rpc.Client
	traceRpcClient *rpc.Client
	proxyRpcClient *rpc.Client

	cache *cache.Cache
}

// NewValidator creates a new inspection validator.
func NewValidator(ctx context.Context, inspectionCfg inspect.InspectionConfig) (*InspectionValidator, error) {
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

type referenceData struct {
	scanApiBlockHash  string
	proxyApiBlockHash string
	traceApiBlockHash string
	traceApiTraceHash string
}

// Validate validates the inspection result. This is intended for using as a first check.
// If this validation fails, nothing else should matter because it is the first and foremost requirement
// to run a node with reliable chain data sources.
func (v *InspectionValidator) Validate(ctx context.Context, results *inspect.InspectionResults) (validationErrs ValidationErrors, resultErr error) {
	refData, err := v.getReferenceData(ctx, results)
	if err != nil {
		validationErrs = validationErrorsFrom(err)
		resultErr = err
		return
	}

	// check if this validator is messing up
	if (refData.scanApiBlockHash != refData.traceApiBlockHash && v.inspectionCfg.CheckTrace) ||
		refData.scanApiBlockHash != refData.proxyApiBlockHash {
		resultErr = multierror.Append(resultErr, ErrReferenceBlockMismatch)
	}

	// check if the inspected node is messing up
	traceApiBlockHashIsDifferent := results.Metadata[inspect.MetadataScanAPIBlockByNumberHash] != results.Metadata[inspect.MetadataTraceAPIBlockByNumberHash] &&
		v.inspectionCfg.CheckTrace

	proxyApiBlockHashIsDifferent := results.Metadata[inspect.MetadataScanAPIBlockByNumberHash] != results.Metadata[inspect.MetadataProxyAPIBlockByNumberHash]

	if traceApiBlockHashIsDifferent || proxyApiBlockHashIsDifferent {
		resultErr = multierror.Append(resultErr, ErrResultBlockMismatch)
	}
	if results.Metadata[inspect.MetadataScanAPIBlockByNumberHash] != refData.scanApiBlockHash {
		resultErr = multierror.Append(resultErr, ErrResultScanAPIBlockMismatch)
	}
	if v.inspectionCfg.CheckTrace {
		if results.Metadata[inspect.MetadataTraceAPIBlockByNumberHash] != refData.traceApiBlockHash {
			resultErr = multierror.Append(resultErr, ErrResultTraceAPIBlockMismatch)
		}
		if results.Metadata[inspect.MetadataTraceAPITraceBlockHash] != refData.traceApiTraceHash {
			resultErr = multierror.Append(resultErr, ErrResultTraceAPITraceBlockMismatch)
		}
	}
	if results.Metadata[inspect.MetadataProxyAPIBlockByNumberHash] != refData.proxyApiBlockHash {
		resultErr = multierror.Append(resultErr, ErrResultProxyAPIBlockMismatch)
	}

	validationErrs = validationErrorsFrom(resultErr)
	return
}

func (v *InspectionValidator) getReferenceData(ctx context.Context, results *inspect.InspectionResults) (refData referenceData, resultErr error) {
	blockNumber := results.Inputs.BlockNumber
	blockNumberStr := strconv.FormatUint(blockNumber, 10)
	if item, ok := v.cache.Get(blockNumberStr); ok {
		return item.(referenceData), nil
	}

	var err error
	refData.scanApiBlockHash, err = inspect.GetBlockResponseHash(ctx, v.scanRpcClient, blockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get scanApiBlockHash api block response hash")
		resultErr = multierror.Append(resultErr, ErrReferenceScanAPIBlock)
	}

	if v.inspectionCfg.CheckTrace {
		refData.traceApiBlockHash, err = inspect.GetBlockResponseHash(ctx, v.traceRpcClient, blockNumber)
		if err != nil {
			log.WithError(err).Error("failed to get trace api block response hash")
			resultErr = multierror.Append(resultErr, ErrReferenceTraceAPIBlock)
		}
		refData.traceApiTraceHash, err = inspect.GetTraceResponseHash(ctx, v.traceRpcClient, blockNumber)
		if err != nil {
			log.WithError(err).Error("failed to get trace api trace block response hash")
			resultErr = multierror.Append(resultErr, ErrReferenceTraceAPITraceBlock)
		}
	}

	refData.proxyApiBlockHash, err = inspect.GetBlockResponseHash(ctx, v.proxyRpcClient, blockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get proxy api block response hash")
		resultErr = multierror.Append(resultErr, ErrReferenceProxyAPIBlock)
	}

	v.cache.Set(blockNumberStr, refData, cacheExpiryDuration)
	return
}

func validationErrorsFrom(err error) (validationErrs ValidationErrors) {
	merr, ok := err.(*multierror.Error)
	if !ok {
		return
	}
	for _, merrErr := range merr.Errors {
		verr, ok := merrErr.(ValidationError)
		if ok {
			validationErrs = append(validationErrs, verr)
		}
	}
	return
}
