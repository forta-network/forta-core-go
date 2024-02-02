package domain

import "testing"

func TestIsMetricAllowed(t *testing.T) {
    tests := []struct {
        name   string
        metric string
        want   bool
    }{
        {
            name:   "Test with allowed metric",
            metric: MetricClientDial,
            want:   true,
        },
        {
            name:   "Test with disallowed metric",
            metric: "agent.client.unknown",
            want:   false,
        },
        {
            name:   "Test with allowed metric prefix",
            metric: "jsonrpc.latency.chainId",
            want:   true,
        },
        {
            name:   "Test with disallowed metric prefix",
            metric: "agent.unknown.",
            want:   false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsMetricAllowed(tt.metric); got != tt.want {
                t.Errorf("IsMetricAllowed() = %v, want %v", got, tt.want)
            }
        })
    }
}