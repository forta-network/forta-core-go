package ipfs

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/forta-protocol/forta-core-go/testutils/testhttp"
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
