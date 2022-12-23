package registeruser

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, status := decodeMessage(message)

	switch status.Value {
	case constants.Ok:
		db.RegisterNewUser(request)

	case constants.BadRequestErr:
		logger.Logger.Sugar().Warn("bad request for registering new user:", message.Value)
	}
}

func decodeMessage(message *kafka.Message) (request *constants.NewUserCreated, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.NewUserCreated) (status *constants.Status) {
	if request.User_id == nil {
		return constants.GenerateBadRequest("user id cannot be null")
	}
	if request.Name == "" {
		return constants.GenerateBadRequest("name cannot be empty")
	}
	return constants.Status_Ok
}
