package registry

import (
	"time"

	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/sirupsen/logrus"
)

const UpdatePaymentSubscription = "UpdatePaymentSubscription"

type UpdatePaymentSubscriptionMessage struct {
	regmsg.Message
	*PaymentSubscription
}

// PaymentSubscription refers to a payment subscription of any type.
type PaymentSubscription struct {
	UserAddress     string `json:"userAddress"`
	Type            string `json:"type"`
	ContractAddress string `json:"contractAddress"`
	TokenID         string `json:"tokenId"`
	Active          bool   `json:"active"`
}

func (m *UpdatePaymentSubscriptionMessage) LogFields() logrus.Fields {
	return logrus.Fields{"type": m.Type, "tokenId": m.TokenID, "userAddress": m.UserAddress}
}

func NewUpdatePaymentSubscriptionMessage(paymentSub *PaymentSubscription) *UpdatePaymentSubscriptionMessage {
	return &UpdatePaymentSubscriptionMessage{
		Message: regmsg.Message{
			Action:    UpdatePaymentSubscription,
			Timestamp: time.Now().UTC(),
		},
		PaymentSubscription: paymentSub,
	}
}
