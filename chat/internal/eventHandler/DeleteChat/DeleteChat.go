package deletechat

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
		// TODO: handle bad room creation request
		log.Fatal("Bad Request")
	}
	res := db.DeleteChat(messagereq)
	if res == nil {
		// TODO: handle errcql while deleting chat
		log.Fatal("Server Error")
	}
	// TODO: return deleted chat status
	log.Println(res)
}

func decodeMessage(message *kafka.Message) (request *constants.MessageDeleteRequest, err error) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, eh.ErrBadRequest
	}
	return request, nil
}
