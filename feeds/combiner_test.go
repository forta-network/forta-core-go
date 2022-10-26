package feeds

import (
	"context"
	"sync"
	"testing"

	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/utils"
)

func Test_combinerFeed_StartRange(t *testing.T) {
	type fields struct {
		start            uint64
		end              uint64
		offset           int
		ctx              context.Context
		cache            utils.Cache
		started          bool
		rateLimit        *time.Ticker
		lastAlert        health.MessageTracker
		alertCh          chan *domain.AlertEvent
		client           *graphql.Client
		botSubscriptions map[string][]string
		botsMu           sync.RWMutex
		alertCache       utils.Cache
		handlers         []cfHandler
	}
	type args struct {
		start uint64
		end   uint64
		rate  int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				cf := &combinerFeed{
					start:            tt.fields.start,
					end:              tt.fields.end,
					offset:           tt.fields.offset,
					ctx:              tt.fields.ctx,
					cache:            tt.fields.cache,
					started:          tt.fields.started,
					rateLimit:        tt.fields.rateLimit,
					lastAlert:        tt.fields.lastAlert,
					alertCh:          tt.fields.alertCh,
					client:           tt.fields.client,
					botSubscriptions: tt.fields.botSubscriptions,
					botsMu:           tt.fields.botsMu,
					alertCache:       tt.fields.alertCache,
					handlers:         tt.fields.handlers,
				}
				cf.StartRange(tt.args.start, tt.args.end, tt.args.rate)
			},
		)
	}
}
