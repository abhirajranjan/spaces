package newcommunity

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, _status := decodeMessage(message)
	switch _status.Value {
	case status.BadRequestErrCode:
		//TODO: handle bad request error
	case status.OkCode:
		res, _status := db.NewCommunity(request)
		switch _status.Value {
		case status.InternalServerErrCode:
			//TODO: handle server error
		case status.OkCode:
			// TODO: return new community data
			print(res)
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.NewCommunityRequest, s *status.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, status.GenerateBadRequest("poorly formated data")
	}
	s = checkRequestForNecessaryData(request)
	return request, s
}

func checkRequestForNecessaryData(request *constants.NewCommunityRequest) (s *status.Status) {
	if request.Name == "" && request.Tag == "" {
		return status.GenerateBadRequest("name and tag both cannot be empty")
	}
	if request.Display_name == "" {
		return status.GenerateBadRequest("display name cannot be empty")
	}
	var testRequest constants.GetCommunityRequest
	testRequest.Name = request.Name
	testRequest.Tag = request.Tag

	switch _, tempstatus := db.GetCommunity(&testRequest); tempstatus.Value {
	case status.OkCode:
		return status.DataAlreadyExists
	case status.InternalServerErrCode:
		return status.ErrDb
	case status.NoDataFoundCode:
		return status.Ok
	default:
		logger.Logger.Sugar().Warn("unknown error type encountered", tempstatus)
		return status.ErrDb
	}
}
