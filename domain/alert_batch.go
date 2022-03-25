package domain

import (
	"github.com/forta-protocol/forta-core-go/encoding"
	"github.com/forta-protocol/forta-core-go/protocol"
	"github.com/forta-protocol/forta-core-go/security"
	"github.com/golang/protobuf/proto"
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

func verifyAndDecode(p *protocol.SignedPayload, target proto.Message) error {
	if err := security.VerifySignedPayload(p); err != nil {
		return err
	}
	return encoding.DecodeGzippedProto(p.Encoded, target)
}

func (a *AlertBatch) Batch() (*protocol.AlertBatch, error) {
	var res protocol.AlertBatch
	if err := verifyAndDecode(a.SignedBatch, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *AlertBatch) BatchSummary() (*protocol.BatchSummary, error) {
	var res protocol.BatchSummary
	if err := verifyAndDecode(a.SignedBatchSummary, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
