package getcommunity

import (
	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"encoding/json"
)

func Handle(message *kafka.Message) {
	request, _status := decodeMessage(message)
	switch _status.Value {
	case status.BadRequestErr:
		// TODO: handle bad request err
	case status.Ok:
		community, _status := db.GetCommunity(request)
		switch _status.Value {
		case status.InternalServerErr:
			//TODO: handle errdb
			print(community)
		case status.NoDataFound:
			// TODO: handle no data matched
		case status.Ok:
			// TODO: return community json
		default:
			logger.Logger.Sugar().Warn("no error matching pattern found", _status)
			// do same as internal server error
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.GetCommunityRequest, s *status.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, status.GenerateBadRequest("poorly formatted data")
	}
	s = checkRequestForNecessaryData(request)
	return request, s
}

func checkRequestForNecessaryData(request *constants.GetCommunityRequest) (s *status.Status) {
	if request.Name == "" && request.Tag == "" {
		return status.GenerateBadRequest("name and tag both cannot be empty")
	}
	return status.OkStatus
}
