package transform

import (
	"reflect"

	"github.com/forta-network/forta-core-go/protocol"
)

func Equal(a, b *protocol.CombinerBotSubscription) bool {
	if a == nil || b == nil {
		return false
	}

	return reflect.DeepEqual(a, b)
}
