package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/security"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/repo"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Env struct {
	PrivateKey string `envconfig:"PRIVATE_KEY"`
	AlertAPI   string `envconfig:"ALERT_API"`
	BatchFile  string `envconfig:"BATCH_FILE"`
}

var (
	env   Env
	batch protocol.AlertBatch
)

func main() {
	if err := envconfig.Process("", &env); err != nil {
		panic(err)
	}

	_, err := url.Parse(env.AlertAPI)
	if err != nil {
		panic(err)
	}
	privKey, err := crypto.HexToECDSA(env.PrivateKey)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(env.BatchFile)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(b, &batch); err != nil {
		panic(err)
	}

	pubKey := privKey.Public().(*ecdsa.PublicKey)
	addr := crypto.PubkeyToAddress(*pubKey)
	key := &keystore.Key{
		Address:    addr,
		PrivateKey: privKey,
	}

	log.WithField("address", addr.Hex()).Info("signing batch")

	signedBatch, err := security.SignBatch(key, &batch)
	if err != nil {
		panic(err)
	}
	signB, _ := json.Marshal(signedBatch)

	ipfsApi, err := coreapi.NewCoreAPI(&core.IpfsNode{
		Repo: &repo.Mock{},
	})
	if err != nil {
		panic(err)
	}
	path, err := ipfsApi.Unixfs().Add(
		context.Background(),
		files.NewBytesFile(createFileBytes(signB)),
		options.Unixfs.HashOnly(true),
	)
	if err != nil {
		panic(err)
	}
	cid := path.Cid().String()

	log.WithField("cid", cid).Info("calculated ipfs hash")

	signedBatchSummary, err := security.SignBatchSummary(
		key, &protocol.BatchSummary{
			Batch:            cid,
			ChainId:          batch.ChainId,
			BlockStart:       batch.BlockStart,
			BlockEnd:         batch.BlockEnd,
			AlertCount:       batch.AlertCount,
			ScannerVersion:   batch.ScannerVersion,
			LatestBlockInput: batch.LatestBlockInput,
			Timestamp:        time.Now().UTC().Format(time.RFC3339),
		},
	)
	if err != nil {
		panic(err)
	}

	reqBody := &domain.AlertBatchRequest{
		Scanner:            addr.Hex(),
		ChainID:            int64(batch.ChainId),
		BlockStart:         int64(batch.BlockStart),
		BlockEnd:           int64(batch.BlockEnd),
		AlertCount:         int64(batch.AlertCount),
		MaxSeverity:        int64(batch.MaxSeverity),
		Ref:                cid,
		SignedBatch:        signedBatch,
		SignedBatchSummary: signedBatchSummary,
	}
	reqB, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	scannerJwt, err := security.CreateScannerJWT(
		key, map[string]interface{}{
			"batch": cid,
		},
	)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/batch/%s", env.AlertAPI, cid), bytes.NewBuffer(reqB))
	if err != nil {
		panic(err)
	}
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", scannerJwt)}

	doReq(req)
}

func doReq(req *http.Request) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.WithFields(log.Fields{
		"response": string(b),
		"status":   resp.StatusCode,
	}).Info("received response")
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
