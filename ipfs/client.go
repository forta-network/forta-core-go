package ipfs

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

const ipfsTimeout = 10 * time.Second

var ErrInternalErr = errors.New("internal error")
var ErrRateLimit = errors.New("rate limited")
var ErrNotFound = errors.New("not found")

type Client interface {
	GetBytes(ctx context.Context, reference string) ([]byte, error)
	UnmarshalJson(ctx context.Context, reference string, target interface{}) error
}

type client struct {
	ipfsGateway string
}

func (c *client) buildUrl(reference string) string {
	return fmt.Sprintf("%s/ipfs/%s", c.ipfsGateway, reference)
}

func (c *client) UnmarshalJson(ctx context.Context, reference string, target interface{}) error {
	b, err := c.GetBytes(ctx, reference)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}

func (c *client) GetBytes(ctx context.Context, reference string) ([]byte, error) {
	ctx, cncl := context.WithTimeout(ctx, ipfsTimeout)
	defer cncl()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.buildUrl(reference), nil)
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewClient(ipfsGateway string) (*client, error) {
	return &client{ipfsGateway: ipfsGateway}, nil
}
