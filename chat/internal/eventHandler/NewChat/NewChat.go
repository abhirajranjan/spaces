package newchat

import (
	"encoding/json"
	"log"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, status := decodeMessage(message)
	switch status.Value {
	case constants.Ok:
		res, status := db.CreateMessage(request)
		if status.Value != constants.Ok {
			// TODO: handle errcql while creating message
			log.Fatal("Server Error", res)
		}
		// TODO: return message created

	case constants.BadRequestErr:
		// TODO: handle bad request
	}
}

func decodeMessage(message *kafka.Message) (request *constants.MessageRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, &request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.MessageRequest) (status *constants.Status) {
	if request.Author_id == nil {
		return constants.GenerateBadRequest("author id cannot be null")
	}
	if request.Room_id == nil {
		return constants.GenerateBadRequest("room id cannot be null")
	}
	if request.Content == "" {
		return constants.GenerateBadRequest("content cannot be empty")
	}
	return constants.Status_Ok
}
