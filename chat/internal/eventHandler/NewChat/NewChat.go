package newchat

import (
	"encoding/json"
	"log"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	eh "github.com/abhirajranjan/spaces/chat/internal/eventHandler"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(request *kafka.Message) {
	messagereq, err := decodeMessage(request)
	if err == eh.ErrBadRequest {
		// TODO: handle bad message request
		log.Fatal("bad request")
		return
	}
	message := db.CreateMessage(messagereq)
	if message == nil {
		// TODO: handle errcql while creating message
		log.Fatal("Server Error")
	}
	// TODO: return message created
	log.Println(message)
}

func decodeMessage(request *kafka.Message) (message *constants.MessageRequest, err error) {
	if err := json.Unmarshal(request.Value, &message); err != nil {
		return nil, eh.ErrBadRequest
	}
	return message, nil
}
