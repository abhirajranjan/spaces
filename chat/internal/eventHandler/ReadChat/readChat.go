package readchat

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
		res, status := db.ReadMessage(request)
		if status.Value == constants.Ok {
			// TODO: handle errcql while creating message
			log.Fatal("Server Error")
		}
		// TODO: return message created
		log.Println(res)

	case constants.BadRequestErr:
		// TODO: handle bad request
	}

}

func decodeMessage(message *kafka.Message) (request *constants.MessageReadRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.MessageReadRequest) (status *constants.Status) {
	if request.Room_id == nil {
		return constants.GenerateBadRequest("room id cannot be null")
	}
	if request.PageSize == nil {
		return constants.GenerateBadRequest("page size cannot be null")
	}
	if request.Author_id == nil {
		return constants.GenerateBadRequest("author id cannot be null")
	}
	return constants.Status_Ok
}
