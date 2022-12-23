package deleteroom

import (
	"encoding/json"
	"log"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	eh "github.com/abhirajranjan/spaces/chat/internal/eventHandler"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	messagereq, err := decodeMessage(message)
	if err == eh.ErrBadRequest {
		// TODO:  handle
		log.Fatal()
	}
	db.DeleteRoom(messagereq)
}

func decodeMessage(message *kafka.Message) (messagereq *constants.DeleteRoomRequest, err error) {
	if err = json.Unmarshal(message.Value, messagereq); err != nil {
		return nil, eh.ErrBadRequest
	}
	return messagereq, nil
}
