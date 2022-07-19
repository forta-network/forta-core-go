package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/golang/protobuf/jsonpb"

	"github.com/forta-network/forta-core-go/testutils/testhttp"
	"github.com/stretchr/testify/assert"
)

type testObj struct {
	Name string `json:"name"`
}

func TestClient_UnmarshalJson(t *testing.T) {
	s := testhttp.NewHttpServer(testhttp.MockHttpConfig{
		MockAPIs: []testhttp.MockApi{
			{
				Method: "GET",
				Path:   "/ipfs/ref",
				ResponseBody: testObj{
					Name: "test",
				},
			},
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	assert.NoError(t, s.Start(ctx))

	c, err := NewClient(s.ServerURL())
	assert.NoError(t, err)

	var result testObj
	err = c.UnmarshalJson(context.Background(), "ref", &result)
	assert.NoError(t, err)

	assert.Equal(t, "test", result.Name)
}

func TestClient_AddFile(t *testing.T) {
	c, err := NewClient("https://ipfs.forta.network")
	assert.NoError(t, err)
	hash := "QmS8cYdRytw1ZeaNwocxASz6Zk7vc28etxXYXv8azu7U6N"
	b, err := os.ReadFile(fmt.Sprintf("./testfiles/%s", hash))
	assert.NoError(t, err)

	var sp protocol.SignedPayload
	u := jsonpb.Unmarshaler{AllowUnknownFields: true}
	assert.NoError(t, u.Unmarshal(bytes.NewReader(b), &sp))

	jb, err := json.Marshal(&sp)
	assert.NoError(t, err)

	res, err := c.AddFile(jb)
	assert.NoError(t, err)
	assert.Equal(t, hash, res)
}

func TestClient_CalculateHash_FromFile(t *testing.T) {
	c, err := NewClient("https://ipfs.forta.network")
	assert.NoError(t, err)
	hash := "QmS8cYdRytw1ZeaNwocxASz6Zk7vc28etxXYXv8azu7U6N"

	b, err := os.ReadFile(fmt.Sprintf("./testfiles/%s", hash))
	assert.NoError(t, err)

	res, err := c.CalculateFileHash(b)
	assert.NoError(t, err)
	assert.Equal(t, hash, res)
}

func TestClient_CalculateHash_FromJson(t *testing.T) {
	c, err := NewClient("https://ipfs.forta.network")
	assert.NoError(t, err)
	hash := "QmS8cYdRytw1ZeaNwocxASz6Zk7vc28etxXYXv8azu7U6N"

	b, err := os.ReadFile(fmt.Sprintf("./testfiles/%s", hash))
	assert.NoError(t, err)

	var sp protocol.SignedPayload
	u := jsonpb.Unmarshaler{AllowUnknownFields: true}
	assert.NoError(t, u.Unmarshal(bytes.NewReader(b), &sp))

	jb, err := json.Marshal(&sp)
	assert.NoError(t, err)

	res, err := c.CalculateFileHash(jb)
	assert.NoError(t, err)
	assert.Equal(t, hash, res)
}

func TestClient_GetBytes(t *testing.T) {
	s := testhttp.NewHttpServer(testhttp.MockHttpConfig{
		MockAPIs: []testhttp.MockApi{
			{
				Method: "GET",
				Path:   "/ipfs/ref",
				ResponseBody: testObj{
					Name: "test",
				},
			},
		},
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	assert.NoError(t, s.Start(ctx))

	c, err := NewClient(s.ServerURL())
	assert.NoError(t, err)

	var result testObj
	b, err := c.GetBytes(context.Background(), "ref")
	assert.NoError(t, err)

	assert.NoError(t, json.Unmarshal(b, &result))
	assert.Equal(t, "test", result.Name)
}
