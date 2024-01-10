package graphql

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	resp := parseBatchResponse([]byte(testBatchResponse))
	data := (*resp.Data.(*BatchGetAlertsResponse))["alerts0"]

	assert.NotNilf(t, resp, "graphql response can not be nil")
	for i := 0; i < 5; i++ {
		assert.Equal(t, fmt.Sprintf("0x%d", i), data.Alerts[i].Source.SourceAlert.Hash)
		assert.Equal(t, "0xbbb", data.Alerts[i].Source.SourceAlert.BotId)
		assert.Equal(t, "2023-01-01T00:00:00Z", data.Alerts[i].Source.SourceAlert.Timestamp)
		assert.Equal(t, uint64(137), data.Alerts[i].Source.SourceAlert.ChainId)
		assert.Equal(t, "Block height: 17890044", data.Alerts[i].Description)
		assert.Equal(t, uint64(i), data.Alerts[i].Source.Block.Number)
	}
}

func TestGetAlertsBatch(t *testing.T) {
	batchResp := parseBatchResponse([]byte(testBatchResponse))
	responseItem := (*batchResp.Data.(*BatchGetAlertsResponse))["alerts0"]
	expectedAlerts := responseItem.ToAlertEvents()
	tests := []struct {
		desc       string
		inputs     []*AlertsInput
		headers    map[string]string
		setupMock  func(mux *http.ServeMux)
		wantAlerts []*protocol.AlertEvent
		wantErr    bool
	}{
		{
			desc: "Successful Request",
			inputs: []*AlertsInput{
				{
					BlockSortDirection: "ASC",
					CreatedSince:       30,
					First:              3,
				},
			},
			headers: map[string]string{
				"Authorization": "Bearer: token",
			},
			setupMock: func(mux *http.ServeMux) {
				// Here's a simple example of what your setup function might do:
				mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

					fmt.Fprint(w, testBatchResponse)
				})
			},
			wantAlerts: expectedAlerts,
			wantErr:    false,
		},
		{
			desc: "Failure due to server error",
			headers: map[string]string{
				"Authorization": "Bearer: token",
			},
			inputs: []*AlertsInput{
				{
					Bots: []string{"0xabc"},
				},
			},
			setupMock: func(mux *http.ServeMux) {
				// Here's a simple example of what your setup function might do:
				mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "server error", http.StatusInternalServerError)
				})
			},
			wantAlerts: nil,
			wantErr:    true,
		},
		{
			desc: "Failure due to unauthorized",
			inputs: []*AlertsInput{
				{
					Bots: []string{"0xabc"},
				},
			},
			headers: map[string]string{
				"Authorization": "", // No token
			},
			setupMock: func(mux *http.ServeMux) {
				// Here's a simple example of what your setup function might do:
				mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
				})
			},
			wantAlerts: nil,
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			mux := http.NewServeMux()
			ts := httptest.NewUnstartedServer(mux)

			tc.setupMock(mux) // Modify setupMock to accept *http.ServeMux
			ts.Start()
			defer ts.Close()

			// Prepare client
			client := NewClient(fmt.Sprintf("%s/graphql", ts.URL))

			// Get context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Invoke GetAlertsBatch
			gotAlerts, gotErr := client.GetAlertsBatch(ctx, tc.inputs, tc.headers)

			if tc.wantErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			require.Equal(t, tc.wantAlerts, gotAlerts)
		})
	}
}

func TestGetAlerts(t *testing.T) {
	_, resp, err := parseResponse([]byte(testResponse))
	assert.NoError(t, err)
	expectedAlerts := resp.Alerts.ToAlertEvents()
	tests := []struct {
		desc       string
		inputs     *AlertsInput
		headers    map[string]string
		setupMock  func(mux *http.ServeMux)
		wantAlerts []*protocol.AlertEvent
		wantErr    bool
	}{
		{
			desc: "Successful Request",
			inputs: &AlertsInput{
				BlockSortDirection: "ASC",
				CreatedSince:       30,
				First:              3,
			},
			headers: map[string]string{
				"Authorization": "Bearer: token",
			},
			setupMock: func(mux *http.ServeMux) {
				// Here's a simple example of what your setup function might do:
				mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

					fmt.Fprint(w, testResponse)
				})
			},
			wantAlerts: expectedAlerts,
			wantErr:    false,
		},
		{
			desc: "Failure due to server error",
			headers: map[string]string{
				"Authorization": "Bearer: token",
			},
			inputs: &AlertsInput{
				Bots: []string{"0xabc"},
			},
			setupMock: func(mux *http.ServeMux) {
				// Here's a simple example of what your setup function might do:
				mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "server error", http.StatusInternalServerError)
				})
			},
			wantAlerts: nil,
			wantErr:    true,
		},
		{
			desc: "Failure due to unauthorized",
			inputs: &AlertsInput{
				Bots: []string{"0xabc"},
			},
			headers: map[string]string{
				"Authorization": "", // No token
			},
			setupMock: func(mux *http.ServeMux) {
				// Here's a simple example of what your setup function might do:
				mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
				})
			},
			wantAlerts: nil,
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			mux := http.NewServeMux()
			ts := httptest.NewUnstartedServer(mux)

			tc.setupMock(mux) // Modify setupMock to accept *http.ServeMux
			ts.Start()
			defer ts.Close()

			// Prepare client
			client := NewClient(fmt.Sprintf("%s/graphql", ts.URL))

			// Get context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Invoke GetAlertsBatch
			gotAlerts, gotErr := client.GetAlerts(ctx, tc.inputs, tc.headers)

			if tc.wantErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			require.Equal(t, tc.wantAlerts, gotAlerts)
		})
	}
}

const testResponse = `{
  "data": {
    "alerts": {
      "pageInfo": {
        "hasNextPage": false,
        "endCursor": {
          "alertId": "0x0baefe6f0be064d7f3637af75a90964e7c231cb6c35266f51af2ce3539558b93",
          "blockNumber": 17890041
        }
      },
      "alerts": [
        {
          "alertId": "FORTA_5",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:57.888797394Z",
          "description": "Block height: 17890044",
          "hash": "0xaa0afa113af00e52594dcf1d8379a508cf0aec7677ca4e0b595d598285697ab1",
          "metadata": null,
          "name": "[BLOCK] Block divisible by 2",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "CRITICAL",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x55708e28820c86a69a09b6858ee6f2b061dc7bd282651ad43d25522cab6e811b",
              "image": "disco-dev.forta.network/bafybeihpk3dmdc2lj2mnnp2udcnxjqkiwco6ytkyo66yuspec27wzgylsu@sha256:cdffe00e7af3baf5ef6dd3f7681df80401afcd732b6af9d84a4da4265b894006",
              "name": null,
              "reference": "QmQFq7vY4P8Ec4SKGR6SCvSUtKMRWB1CYEuDwWg4oV5MnD",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 0,
              "hash": "0x722436d03af0058ef990e0b6af684a7eb7b4988b4abe90dfb5bbee9b7c93a67b",
              "timestamp": "2023-08-11T07:06:35Z",
              "chainId": 1
            },
            "sourceAlert": {
              "hash": "0x0",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "EXPLOIT",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": [
            {
              "label": "label divisible to 15",
              "confidence": 1,
              "entity": "0x722436d03af0058ef990e0b6af684a7eb7b4988b4abe90dfb5bbee9b7c93a67b",
              "entityType": "UNKNOWN_ENTITY_TYPE",
              "remove": false,
              "metadata": [
                "foo=bar"
              ]
            }
          ]
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:57.820411096Z",
          "description": "Block height: 17890044",
          "hash": "0x073161be18e8deac6a40b3e3bd4a54b984a4b3042763a7cabc16137f49e0d3f4",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x49daccfa279e21f183e46e8a6f33ba6b301ba9e222f7f8aa4f2ec596c214b807",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmeiXnZHw8pjzJzqvcvBynwjvtpKcoHPmoWSUq9kK1PgSu",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 1,
              "hash": "0x4d78ebd97813fd5118c6629313ddc7e8c9204e064d1ce4a725f6e8e740b28ca5",
              "timestamp": "2023-08-11T07:06:23Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x1",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:49.385274239Z",
          "description": "Block height: 17890044",
          "hash": "0x9203279e8a53e3aee0b14649e6b97c7851fca2029d67eb91e9282db8bd32bac2",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x49daccfa279e21f183e46e8a6f33ba6b301ba9e222f7f8aa4f2ec596c214b807",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmeiXnZHw8pjzJzqvcvBynwjvtpKcoHPmoWSUq9kK1PgSu",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 2,
              "hash": "0x8963c2c1acc0103f4b057cebea81dd332d0d98261ffa2bd808a15426b3e4c56e",
              "timestamp": "2023-08-11T07:06:11Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x2",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:29.507245254Z",
          "description": "Block height: 17890044",
          "hash": "0xffed0ab5d547b5d04b9655b343c2a3ed3c0138cf944395291207363b4c43927d",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x49daccfa279e21f183e46e8a6f33ba6b301ba9e222f7f8aa4f2ec596c214b807",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmeiXnZHw8pjzJzqvcvBynwjvtpKcoHPmoWSUq9kK1PgSu",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 3,
              "hash": "0x957b15c7d26e7079b361b877343c8999a1c8cea472469eb95d3aa669bab61669",
              "timestamp": "2023-08-11T07:05:59Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x3",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:29.496884512Z",
          "description": "Block height: 17890044",
          "hash": "0x0baefe6f0be064d7f3637af75a90964e7c231cb6c35266f51af2ce3539558b93",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0xbe1872858e63b6ed4ef7b84fc453970dc8d89968715797662a4f43c01d598aab",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmVUdDykRvMTC5JgxJp3UJrYZfURH3KKrNm5qiiojfekbM",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 4,
              "hash": "0x957b15c7d26e7079b361b877343c8999a1c8cea472469eb95d3aa669bab61669",
              "timestamp": "2023-08-11T07:05:59Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x4",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        }
      ]
    }
  }
}`

const testBatchResponse = `{
  "data": {
    "alerts0": {
      "pageInfo": {
        "hasNextPage": false,
        "endCursor": {
          "alertId": "0x0baefe6f0be064d7f3637af75a90964e7c231cb6c35266f51af2ce3539558b93",
          "blockNumber": 17890041
        }
      },
      "alerts": [
        {
          "alertId": "FORTA_5",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:57.888797394Z",
          "description": "Block height: 17890044",
          "hash": "0xaa0afa113af00e52594dcf1d8379a508cf0aec7677ca4e0b595d598285697ab1",
          "metadata": null,
          "name": "[BLOCK] Block divisible by 2",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "CRITICAL",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x55708e28820c86a69a09b6858ee6f2b061dc7bd282651ad43d25522cab6e811b",
              "image": "disco-dev.forta.network/bafybeihpk3dmdc2lj2mnnp2udcnxjqkiwco6ytkyo66yuspec27wzgylsu@sha256:cdffe00e7af3baf5ef6dd3f7681df80401afcd732b6af9d84a4da4265b894006",
              "name": null,
              "reference": "QmQFq7vY4P8Ec4SKGR6SCvSUtKMRWB1CYEuDwWg4oV5MnD",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 0,
              "hash": "0x722436d03af0058ef990e0b6af684a7eb7b4988b4abe90dfb5bbee9b7c93a67b",
              "timestamp": "2023-08-11T07:06:35Z",
              "chainId": 1
            },
            "sourceAlert": {
              "hash": "0x0",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "EXPLOIT",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": [
            {
              "label": "label divisible to 15",
              "confidence": 1,
              "entity": "0x722436d03af0058ef990e0b6af684a7eb7b4988b4abe90dfb5bbee9b7c93a67b",
              "entityType": "UNKNOWN_ENTITY_TYPE",
              "remove": false,
              "metadata": [
                "foo=bar"
              ]
            }
          ]
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:57.820411096Z",
          "description": "Block height: 17890044",
          "hash": "0x073161be18e8deac6a40b3e3bd4a54b984a4b3042763a7cabc16137f49e0d3f4",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x49daccfa279e21f183e46e8a6f33ba6b301ba9e222f7f8aa4f2ec596c214b807",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmeiXnZHw8pjzJzqvcvBynwjvtpKcoHPmoWSUq9kK1PgSu",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 1,
              "hash": "0x4d78ebd97813fd5118c6629313ddc7e8c9204e064d1ce4a725f6e8e740b28ca5",
              "timestamp": "2023-08-11T07:06:23Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x1",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:49.385274239Z",
          "description": "Block height: 17890044",
          "hash": "0x9203279e8a53e3aee0b14649e6b97c7851fca2029d67eb91e9282db8bd32bac2",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x49daccfa279e21f183e46e8a6f33ba6b301ba9e222f7f8aa4f2ec596c214b807",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmeiXnZHw8pjzJzqvcvBynwjvtpKcoHPmoWSUq9kK1PgSu",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 2,
              "hash": "0x8963c2c1acc0103f4b057cebea81dd332d0d98261ffa2bd808a15426b3e4c56e",
              "timestamp": "2023-08-11T07:06:11Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x2",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:29.507245254Z",
          "description": "Block height: 17890044",
          "hash": "0xffed0ab5d547b5d04b9655b343c2a3ed3c0138cf944395291207363b4c43927d",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0x49daccfa279e21f183e46e8a6f33ba6b301ba9e222f7f8aa4f2ec596c214b807",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmeiXnZHw8pjzJzqvcvBynwjvtpKcoHPmoWSUq9kK1PgSu",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 3,
              "hash": "0x957b15c7d26e7079b361b877343c8999a1c8cea472469eb95d3aa669bab61669",
              "timestamp": "2023-08-11T07:05:59Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x3",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        },
        {
          "alertId": "FORTA-1",
          "addresses": null,
          "contracts": null,
          "createdAt": "2023-08-11T07:06:29.496884512Z",
          "description": "Block height: 17890044",
          "hash": "0x0baefe6f0be064d7f3637af75a90964e7c231cb6c35266f51af2ce3539558b93",
          "metadata": null,
          "name": "src: 0x3f88c2b3e267e6b8e9dE017cdB47a59aC9Ecb284",
          "projects": null,
          "protocol": "ethereum",
          "scanNodeCount": 1,
          "severity": "LOW",
          "source": {
            "transactionHash": "",
            "bot": {
              "chainIds": null,
              "createdAt": null,
              "description": null,
              "developer": "",
              "docReference": null,
              "enabled": null,
              "id": "0xbe1872858e63b6ed4ef7b84fc453970dc8d89968715797662a4f43c01d598aab",
              "image": "disco-dev.forta.network/bafybeianf2tl4rtmd7yabfbwhapllkscnnnjq4lzm7wwqs3jxnam53sbgm@sha256:d210b7e72671eda254fdc6bdf65ecbf6f6fb43c98ebb81523a02d7f9e6ce9c02",
              "name": null,
              "reference": "QmVUdDykRvMTC5JgxJp3UJrYZfURH3KKrNm5qiiojfekbM",
              "repository": null,
              "projects": null,
              "scanNodes": null,
              "version": null
            },
            "block": {
              "number": 4,
              "hash": "0x957b15c7d26e7079b361b877343c8999a1c8cea472469eb95d3aa669bab61669",
              "timestamp": "2023-08-11T07:05:59Z",
              "chainId": 1
            },
           "sourceAlert": {
              "hash": "0x4",
              "botId": "0xbbb",
              "timestamp": "2023-01-01T00:00:00Z",
              "chainId": 137
            }
          },
          "alertDocumentType": "BLOCK",
          "findingType": "INFORMATION",
          "relatedAlerts": null,
          "chainId": 1,
          "labels": []
        }
      ]
    }
  }
}`
