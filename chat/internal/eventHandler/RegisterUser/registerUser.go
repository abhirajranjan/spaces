package registeruser

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	eh "github.com/abhirajranjan/spaces/chat/internal/eventHandler"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	messagereq, err := decodeMessage(message)
	if err == eh.ErrBadRequest {
		logger.Logger.Sugar().Warn("bad request for registering new user:", message.Value)
		return
	}
	db.RegisterNewUser(messagereq)
}

func decodeMessage(message *kafka.Message) (MessageRequest *constants.NewUserCreated, err error) {
	if err = json.Unmarshal(message.Value, MessageRequest); err != nil {
		return nil, eh.ErrBadRequest
	}
	return MessageRequest, nil
}
