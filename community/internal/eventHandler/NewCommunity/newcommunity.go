package newcommunity

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, status := decodeMessage(message)
	switch status.Value {
	case constants.BadRequestErr:
		//TODO: handle bad request error
	case constants.Ok:
		res, status := db.NewCommunity(request)
		switch status.Value {
		case constants.InternalServerErr:
			//TODO: handle server error
		case constants.Ok:
			// TODO: return new community data
			print(res)
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.NewCommunityRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formated data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.NewCommunityRequest) (status *constants.Status) {
	if request.Name == "" && request.Tag == "" {
		return constants.GenerateBadRequest("name and tag both cannot be empty")
	}
	if request.Display_name == "" {
		return constants.GenerateBadRequest("display name cannot be empty")
	}
	var testRequest constants.GetCommunityRequest
	testRequest.Name = request.Name
	testRequest.Tag = request.Tag

	switch _, tempstatus := db.GetCommunity(&testRequest); tempstatus.Value {
	case constants.Ok:
		return constants.Status_AccountAlreadyExists
	case constants.InternalServerErr:
		return constants.Status_ErrDb
	case constants.NoDataFound:
		return constants.Status_Ok
	default:
		logger.Logger.Sugar().Warn("unknown error type encountered", tempstatus)
		return constants.Status_ErrDb
	}
}
