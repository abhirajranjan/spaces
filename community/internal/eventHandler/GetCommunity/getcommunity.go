package getcommunity

import (
	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"encoding/json"
)

func Handle(message *kafka.Message) {
	request, status := decodeMessage(message)
	switch status.Value {
	case constants.BadRequestErr:
		// TODO: handle bad request err
	case constants.Ok:
		community, status := db.GetCommunity(request)
		switch status.Value {
		case constants.InternalServerErr:
			//TODO: handle errdb
			print(community)
		case constants.NoDataFound:
			// TODO: handle no data matched
		case constants.Ok:
			// TODO: return community json
		default:
			logger.Logger.Sugar().Warn("no error matching pattern found", status)
			// do same as internal server error
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.GetCommunityRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.GetCommunityRequest) (status *constants.Status) {
	if request.Name == "" && request.Tag == "" {
		return constants.GenerateBadRequest("name and tag both cannot be empty")
	}
	return constants.Status_Ok
}
