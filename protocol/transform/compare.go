package transform

import "github.com/forta-network/forta-core-go/protocol"

func CompareCombinerBotSubscription(a, b *protocol.CombinerBotSubscription) bool {
	if a == nil || b == nil {
		return false
	}

	sameBot := a.BotId == b.BotId
	sameAlert := a.AlertId == b.AlertId

	return sameAlert && sameBot
}
