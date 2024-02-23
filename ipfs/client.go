package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/forta-network/forta-core-go/utils/httpclient"
	ipfsapi "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-libipfs/files"
	config "github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/ipfs/kubo/repo"

	log "github.com/sirupsen/logrus"
)

const ipfsTimeout = 10 * time.Second

var ErrInternalErr = errors.New("internal error")
var ErrRateLimit = errors.New("rate limited")
var ErrNotFound = errors.New("not found")

type Client interface {
	AddFile(payload []byte) (string, error)
	CalculateFileHash(payload []byte) (string, error)
	GetBytes(ctx context.Context, reference string) ([]byte, error)
	UnmarshalJson(ctx context.Context, reference string, target interface{}) error
}

type client struct {
	ipfsGateway string
	coreApi     coreiface.CoreAPI
	*ipfsapi.Shell
}

func (c *client) buildUrl(reference string) string {
	return fmt.Sprintf("%s/ipfs/%s", c.ipfsGateway, reference)
}

func createFileBytes(payload []byte) []byte {
	str := string(payload)
	// if written to file and uploaded to ipfs, final character would be a \n
	// this simulates that action so that hashes are same
	if !strings.HasSuffix(str, "\n") {
		str = fmt.Sprintf("%s%s", str, "\n")
	}
	return []byte(str)
}

func (c *client) AddFile(payload []byte) (string, error) {
	b := createFileBytes(payload)
	return c.Add(bytes.NewReader(b), ipfsapi.Pin(true))
}

func (c *client) CalculateFileHash(payload []byte) (string, error) {
	b := createFileBytes(payload)
	path, err := c.coreApi.Unixfs().Add(
		context.Background(),
		files.NewBytesFile(b),
		options.Unixfs.HashOnly(true),
	)
	if err != nil {
		return "", err
	}

	splittedPath := strings.Split(path.String(), "/")
	return splittedPath[len(splittedPath)-1], nil
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
	resp, err := httpclient.Default.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
	coreApi, err := coreapi.NewCoreAPI(&core.IpfsNode{
		Repo: &repo.Mock{
			C: config.Config{},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create core api: %v", err)
	}

	return &client{
		ipfsGateway: ipfsGateway,
		coreApi:     coreApi,
		Shell:       ipfsapi.NewShell(ipfsGateway),
	}, nil
}
