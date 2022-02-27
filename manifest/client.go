package manifest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const manifestTimeout = 10 * time.Second

var ErrManifestParse = errors.New("json parse error")
var ErrInternalErr = errors.New("internal error")
var ErrRateLimit = errors.New("rate limited")
var ErrNotFound = errors.New("not found")

type Client interface {
	GetAgentManifest(ctx context.Context, reference string) (*SignedAgentManifest, error)
}

type manifestStore struct {
	ipfsGateway string
}

func (ms *manifestStore) buildUrl(reference string) string {
	return fmt.Sprintf("%s/ipfs/%s", ms.ipfsGateway, reference)
}

func (ms *manifestStore) GetManifest(ctx context.Context, reference string) (*SignedAgentManifest, error) {
	ctx, cncl := context.WithTimeout(ctx, manifestTimeout)
	defer cncl()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, ms.buildUrl(reference), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 500 {
		log.WithFields(log.Fields{
			"reference": reference,
		}).Errorf("5xx error retrieving from ipfs: %d", resp.StatusCode)
		return nil, ErrInternalErr
	}
	if resp.StatusCode == 429 {
		log.WithFields(log.Fields{
			"reference": reference,
		}).Errorf("rate limited retrieving from ipfs: %d", resp.StatusCode)
		return nil, ErrRateLimit
	}

	// ipfs will return a 400 if an invalid ipfs, which to us is a not found
	if resp.StatusCode >= 400 {
		log.WithFields(log.Fields{
			"reference": reference,
		}).Warnf("4xx error retrieving from ipfs: %d", resp.StatusCode)
		return nil, ErrNotFound
	}

	var m SignedAgentManifest
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, ErrManifestParse
	}
	return &m, nil
}

func NewClient(ipfsGateway string) (*manifestStore, error) {
	return &manifestStore{ipfsGateway: ipfsGateway}, nil
}
