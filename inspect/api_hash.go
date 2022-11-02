package inspect

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
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

// HashReferences contains hash references that are used during inspection validation.
type HashReferences struct {
	ScanAPIBlockHash  string
	ProxyAPIBlockHash string
	TraceAPIBlockHash string
	TraceAPITraceHash string
}

// ValidateHashReferences validates results against the references.
func ValidateHashReferences(metadata map[string]string, cfg InspectionConfig, refData *HashReferences) (validationErrs ValidationErrors, resultErr error) {
	// check if this validator is messing up
	if (refData.ScanAPIBlockHash != refData.TraceAPIBlockHash && cfg.CheckTrace) ||
		refData.ScanAPIBlockHash != refData.ProxyAPIBlockHash {
		resultErr = multierror.Append(resultErr, ErrReferenceBlockMismatch)
	}

	// check if the inspected node is messing up
	traceApiBlockHashIsDifferent := metadata[MetadataScanAPIBlockByNumberHash] != metadata[MetadataTraceAPIBlockByNumberHash] &&
		cfg.CheckTrace

	proxyApiBlockHashIsDifferent := metadata[MetadataScanAPIBlockByNumberHash] != metadata[MetadataProxyAPIBlockByNumberHash]

	if traceApiBlockHashIsDifferent || proxyApiBlockHashIsDifferent {
		resultErr = multierror.Append(resultErr, ErrResultBlockMismatch)
	}
	if metadata[MetadataScanAPIBlockByNumberHash] != refData.ScanAPIBlockHash {
		resultErr = multierror.Append(resultErr, ErrResultScanAPIBlockMismatch)
	}
	if cfg.CheckTrace {
		if metadata[MetadataTraceAPIBlockByNumberHash] != refData.TraceAPIBlockHash {
			resultErr = multierror.Append(resultErr, ErrResultTraceAPIBlockMismatch)
		}
		if metadata[MetadataTraceAPITraceBlockHash] != refData.TraceAPITraceHash {
			resultErr = multierror.Append(resultErr, ErrResultTraceAPITraceBlockMismatch)
		}
	}
	if metadata[MetadataProxyAPIBlockByNumberHash] != refData.ProxyAPIBlockHash {
		resultErr = multierror.Append(resultErr, ErrResultProxyAPIBlockMismatch)
	}

	validationErrs = ValidationErrorsFrom(resultErr)
	return
}

// ValidationErrorsFrom extracts the validation errors from given error.
func ValidationErrorsFrom(err error) (validationErrs ValidationErrors) {
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
