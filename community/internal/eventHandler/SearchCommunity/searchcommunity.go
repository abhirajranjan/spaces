package searchcommunity

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, status := decodeMessage(message)
	switch status.Value {
	case constants.BadRequestErr:
		// TODO: handle bad requests
	case constants.Ok:
		searchresponse, status := db.SearchCommunity(request)
		if status.Value != constants.Ok {
			// TODO: handle errdb
			print(searchresponse)
		}
		// TODO handle search response
		print(searchresponse)
	}
}

func decodeMessage(message *kafka.Message) (request *constants.SearchCommunityRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	status = checkRequestForNecessaryData(request)
	return request, status
}

func checkRequestForNecessaryData(request *constants.SearchCommunityRequest) (status *constants.Status) {
	if request.Name == "" && request.Tag == "" {
		return constants.GenerateBadRequest("name and tag both cannot be empty")
	}
	if request.Pagesize == 0 {
		return constants.GenerateBadRequest("page size cannot be null")
	}
	return constants.Status_Ok
}
