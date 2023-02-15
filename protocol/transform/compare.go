package transform

import (
	"sort"
	"strings"

	"github.com/forta-network/forta-core-go/protocol"
)

func Equal(a, b *protocol.CombinerBotSubscription) bool {
	if a == nil || b == nil {
		return false
	}

	if a.BotId != b.BotId {
		return false
	}

	if a.AlertId != b.AlertId {
		return false
	}

	if a.ChainId != b.ChainId {
		return false
	}

	sort.Strings(a.AlertIds)
	sort.Strings(b.AlertIds)

	return strings.EqualFold(strings.Join(a.AlertIds, ","), strings.Join(b.AlertIds, ","))
}
