package inspect

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/hashicorp/go-multierror"
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
	ErrReferenceBlockMismatch = &validationError{
		ErrorCode:    "2011",
		ErrorMessage: "reference scan and trace apis return different blocks",
	}

	// 3xxx: node messed up: inspection results are bad
	// 3x0x: inspection results have bad/mismatching data
	// 3x1x: inspection results do not match with the reference
	ErrResultBlockMismatch = &validationError{
		ErrorCode:    "3001",
		ErrorMessage: "scan and trace inspection detected different blocks",
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
)

// InspectionValidator validates inspection results.
type InspectionValidator struct {
	inspectionCfg  *InspectionConfig
	scanRpcClient  *rpc.Client
	traceRpcClient *rpc.Client
	traceClient    ethereum.Client
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
	validator.inspectionCfg = &inspectionCfg
	return &validator, nil
}

// Validate validates the inspection result. This is intended for using as a first check.
// If this validation fails, nothing else should matter because it is the first and foremost requirement
// to run a node with reliable chain data sources.
func (v *InspectionValidator) Validate(ctx context.Context, results *InspectionResults) (resultErr error) {
	scanApiBlockResponseHash, err := getBlockResponseHash(ctx, v.scanRpcClient, v.inspectionCfg.BlockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get scan api block response hash")
		resultErr = multierror.Append(resultErr, ErrReferenceScanAPI)
	}
	traceApiBlockResponseHash, err := getBlockResponseHash(ctx, v.traceRpcClient, v.inspectionCfg.BlockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get trace api block response hash")
		resultErr = multierror.Append(resultErr, ErrReferenceTraceAPI)
	}
	traceApiTraceBlockResponseHash, err := getTraceResponseHash(ctx, v.traceClient, v.inspectionCfg.BlockNumber)
	if err != nil {
		log.WithError(err).Error("failed to get trace api trace block response hash")
		resultErr = multierror.Append(resultErr, ErrReferenceTraceAPI)
	}

	// return with the first batch of errors at this step
	if resultErr != nil {
		return
	}

	if results.Metadata[MetadataScanAPIBlockByNumberHash] != results.Metadata[MetadataTraceAPIBlockByNumberHash] {
		resultErr = multierror.Append(resultErr, ErrResultBlockMismatch)
	}
	if results.Metadata[MetadataScanAPIBlockByNumberHash] != scanApiBlockResponseHash {
		resultErr = multierror.Append(resultErr, ErrResultScanAPIBlockMismatch)
	}
	if results.Metadata[MetadataTraceAPIBlockByNumberHash] != traceApiBlockResponseHash {
		resultErr = multierror.Append(resultErr, ErrResultTraceAPIBlockMismatch)
	}
	if results.Metadata[MetadataTraceAPITraceBlockHash] != traceApiTraceBlockResponseHash {
		resultErr = multierror.Append(resultErr, ErrResultTraceAPITraceBlockMismatch)
	}

	return
}
