package createroom

import (
	"encoding/json"
	"log"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, err := decodeMessage(message)
	switch err.Value {
	case constants.Ok:
		room, status := db.CreateRoom(request)
		if status.Value != constants.Ok {
			// TODO: handle errcql while creating room
			log.Fatal("Server Error")
		}
		// TODO: return room created
		log.Println(room)

	case constants.BadRequestErr:
		// TODO: handle bad requests
	}
}

func decodeMessage(message *kafka.Message) (request *constants.RoomCreationRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.RoomCreationRequest) (status *constants.Status) {
	// ? description not compulsary ?
	if request.Name == "" {
		return constants.GenerateBadRequest("name cannot be empty")
	}
	if request.Author_id == nil {
		return constants.GenerateBadRequest("author id cannot be null")
	}
	return constants.Status_Ok
}
