package transform

import (
	"testing"

	"github.com/forta-network/forta-core-go/protocol"
)

func TestIsSameBotSubscription(t *testing.T) {
	type args struct {
		a *protocol.CombinerBotSubscription
		b *protocol.CombinerBotSubscription
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "same subscriptions",
			args: args{
				a: &protocol.CombinerBotSubscription{
					BotId: "a", AlertId: "b", ChainId: 1, AlertIds: []string{"a", "b"},
				},
				b: &protocol.CombinerBotSubscription{
					BotId: "a", AlertId: "b", ChainId: 1, AlertIds: []string{"a", "b"},
				},
			},
			want: true,
		},
		{
			name: "same subscription, different chain",
			args: args{
				a: &protocol.CombinerBotSubscription{BotId: "a", AlertId: "b", ChainId: 2},
				b: &protocol.CombinerBotSubscription{BotId: "a", AlertId: "b", ChainId: 1},
			},
			want: false,
		},
		{
			name: "different alert ids",
			args: args{
				a: &protocol.CombinerBotSubscription{BotId: "a", AlertIds: []string{"a", "b"}, ChainId: 2},
				b: &protocol.CombinerBotSubscription{BotId: "a", AlertIds: []string{"a", "x"}, ChainId: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Equal(tt.args.a, tt.args.b); got != tt.want {
					t.Errorf("Equal() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
