package deleteroom

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
		status := db.DeleteRoom(request)
		if status.Value != constants.Ok {
			// TODO: handle errcql while deleting room
			log.Fatal("Server Error")
		}
		// TODO: return deleted room status
	case constants.BadRequestErr:
		// TODO: handle bad request
	}
}

func decodeMessage(message *kafka.Message) (request *constants.DeleteRoomRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.DeleteRoomRequest) (status *constants.Status) {
	if request.Author_id == nil {
		return constants.GenerateBadRequest("author id cannot be null")
	}
	if request.Room_id == nil {
		return constants.GenerateBadRequest("room id cannot be null")
	}
	return constants.Status_Ok
}
