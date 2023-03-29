package graphql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertClient_GetAlerts(t *testing.T) {
	type fields struct {
		url string
	}
	type args struct {
		input *AlertsInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "get alerts",
			args: args{
				input: &AlertsInput{
					Bots: []string{"0x5e13c2f3a97c292695b598090056ba5d52f9dcc7790bcdaa8b6cd87c1a1ebc0f"},
				},
			},
			fields: fields{url: "https://api.forta.network/graphql"},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return err == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ac := NewClient(tt.fields.url)

				ctx := context.Background()
				got, err := ac.GetAlerts(ctx, tt.args.input, nil)
				if !tt.wantErr(t, err, "GetAlerts()") {
					return
				}
				_ = got
			},
		)
	}
}
