package domain

import (
	"github.com/forta-protocol/forta-core-go/encoding"
	"github.com/forta-protocol/forta-core-go/protocol"
	"github.com/forta-protocol/forta-core-go/security"
)

type AlertBatch struct {
	Scanner            string                  `json:"scanner"`
	ChainID            int64                   `json:"chainId"`
	BlockStart         int64                   `json:"blockStart"`
	BlockEnd           int64                   `json:"blockEnd"`
	AlertCount         int64                   `json:"alertCount"`
	MaxSeverity        int64                   `json:"maxSeverity"`
	Ref                string                  `json:"ref"`
	SignedBatch        *protocol.SignedPayload `json:"signedBatch"`
	SignedBatchSummary *protocol.SignedPayload `json:"batchSummary"`
}

func (a *AlertBatch) Batch() (*protocol.AlertBatch, error) {
	if err := security.VerifySignedPayload(a.SignedBatch); err != nil {
		return nil, err
	}
	var res protocol.AlertBatch
	if err := encoding.DecodeGzippedProto(a.SignedBatch.Encoded, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *AlertBatch) BatchSummary() (*protocol.BatchSummary, error) {
	if err := security.VerifySignedPayload(a.SignedBatchSummary); err != nil {
		return nil, err
	}
	var res protocol.BatchSummary
	if err := encoding.DecodeGzippedProto(a.SignedBatchSummary.Encoded, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
