package readchat

import (
	"encoding/json"
	"log"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	eh "github.com/abhirajranjan/spaces/chat/internal/eventHandler"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(req *kafka.Message) {
	messagereq, err := decodeMessage(req)
	if err == eh.ErrBadRequest {
		// TODO: handle bad message request
		log.Fatal("bad request")
		return
	}
	message := db.ReadMessage(messagereq)
	if message == nil {
		// TODO: handle errcql while creating message
		log.Fatal("Server Error")
	}
	// TODO: return message created
	log.Println(message)
}

func decodeMessage(message *kafka.Message) (MessageRequest *constants.MessageReadRequest, err error) {
	if err = json.Unmarshal(message.Value, MessageRequest); err != nil {
		return nil, eh.ErrBadRequest
	}
	return MessageRequest, nil
}
