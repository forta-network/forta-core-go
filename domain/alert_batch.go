package domain

import (
	"github.com/forta-network/forta-core-go/protocol"
)

// AlertBatchRequest is the request payload to the alerts API
type AlertBatchRequest struct {
	Scanner            string                  `json:"scanner"`
	ChainID            int64                   `json:"chainId"`
	BlockStart         int64                   `json:"blockStart"`
	BlockEnd           int64                   `json:"blockEnd"`
	AlertCount         int64                   `json:"alertCount"`
	MaxSeverity        int64                   `json:"maxSeverity"`
	Ref                string                  `json:"ref"`
	SignedBatch        *protocol.SignedPayload `json:"signedBatch"`
	SignedBatchSummary *protocol.SignedPayload `json:"signedBatchSummary"`
}

// AlertBatchResponse is the response from the alerts API
type AlertBatchResponse struct {
	ReceiptID     string                  `json:"receiptId"`
	SignedReceipt *protocol.SignedPayload `json:"signedReceipt"`
}

// AlertBatchMessage is passed to batch processing asynchronously
type AlertBatchMessage struct {
	Scanner       string                  `json:"scanner"`
	ChainID       int64                   `json:"chainId"`
	BlockStart    int64                   `json:"blockStart"`
	BlockEnd      int64                   `json:"blockEnd"`
	AlertCount    int64                   `json:"alertCount"`
	MaxSeverity   int64                   `json:"maxSeverity"`
	Ref           string                  `json:"ref"`
	SignedReceipt *protocol.SignedPayload `json:"signedReceipt"`
}
