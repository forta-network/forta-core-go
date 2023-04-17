package registry

import (
	"github.com/forta-network/forta-core-go/contracts/merged/contract_rewards_distributor"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type RewardedMessage struct {
	regmsg.Message
	Amount      string `json:"amount"`
	Subject     string `json:"subject"`
	SubjectType int    `json:"subjectType"`
	Epoch       int64  `json:"epoch"`
}

func (rm *RewardedMessage) LogFields() logrus.Fields {
	return logrus.Fields{
		"subject": rm.Subject,
		"amount":  rm.Amount,
		"epoch":   rm.Epoch,
	}
}

func NewRewardedMessage(evt *contract_rewards_distributor.RewardsDistributorRewarded, blk *domain.Block) *RewardedMessage {
	return &RewardedMessage{
		Message: regmsg.Message{
			Action:    Rewarded,
			Timestamp: time.Now().UTC(),
			Source:    regmsg.SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		Subject:     evt.Subject.Text(10),
		Amount:      evt.Amount.Text(10),
		SubjectType: int(evt.SubjectType),
		Epoch:       evt.EpochNumber.Int64(),
	}
}

type ClaimedRewardsMessage struct {
	regmsg.Message
	To          string `json:"to"`
	Amount      string `json:"amount"`
	Subject     string `json:"subject"`
	SubjectType int    `json:"subjectType"`
	Epoch       int64  `json:"epoch"`
}

func (cr *ClaimedRewardsMessage) LogFields() logrus.Fields {
	return logrus.Fields{
		"subject": cr.Subject,
		"amount":  cr.Amount,
		"epoch":   cr.Epoch,
		"to":      cr.To,
	}
}

func NewClaimedRewardsMessage(evt *contract_rewards_distributor.RewardsDistributorClaimedRewards, blk *domain.Block) *ClaimedRewardsMessage {
	return &ClaimedRewardsMessage{
		Message: regmsg.Message{
			Action:    ClaimedRewards,
			Timestamp: time.Now().UTC(),
			Source:    regmsg.SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		To:          strings.ToLower(evt.To.Hex()),
		Subject:     evt.Subject.Text(10),
		Amount:      evt.Value.Text(10),
		SubjectType: int(evt.SubjectType),
		Epoch:       evt.EpochNumber.Int64(),
	}
}
