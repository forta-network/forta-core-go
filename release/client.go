package release

import (
	"encoding/json"
	"fmt"
	"github.com/forta-network/forta-core-go/utils/httpclient"
	log "github.com/sirupsen/logrus"
)

type Client interface {
	GetReleaseManifest(reference string) (*ReleaseManifest, error)
}

type client struct {
	prefixes []string
}

func get(url string) (*ReleaseManifest, error) {
	res, err := httpclient.Default.Get(url)
	if err != nil {
		return nil, err
	}
	var rm ReleaseManifest
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&rm); err != nil {
		return nil, err
	}

	if rm.Release.Version == "" {
		return nil, fmt.Errorf("%s: version not set", url)
	}

	return &rm, nil
}

func (c *client) GetReleaseManifest(reference string) (*ReleaseManifest, error) {
	var resErr error
	for _, p := range c.prefixes {
		url := fmt.Sprintf("%s/%s", p, reference)
		logger := log.WithFields(log.Fields{
			"url": url,
		})
		res, err := get(url)
		if err != nil {
			logger.WithError(err).Error("error loading from url")
			resErr = err
			continue
		}
		return res, nil
	}
	return nil, resErr
}

func NewClient(ipfsGateway string, urlPrefixes []string) (Client, error) {
	return &client{prefixes: append(urlPrefixes, fmt.Sprintf("%s/ipfs", ipfsGateway))}, nil
}
