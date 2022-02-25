package registry

import (
	"errors"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Action string `json:"action"`
}

func ParseMessage(msg string) (*Message, error) {
	var r Message
	err := json.Unmarshal([]byte(msg), &r)
	if err != nil {
		return nil, err
	}
	if r.Action == "" {
		log.WithFields(log.Fields{
			"body": msg,
		}).Error("action is not populated")
		return nil, errors.New("action is not populated")
	}
	return &r, nil
}
