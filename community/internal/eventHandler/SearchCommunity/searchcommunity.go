package searchcommunity

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, _status := decodeMessage(message)
	switch _status.Value {
	case status.BadRequestErr:
		// TODO: handle bad requests
	case status.Ok:
		searchresponse, _status := db.SearchCommunity(request)
		if _status.Value != status.Ok {
			// TODO: handle errdb
			print(searchresponse)
		}
		// TODO handle search response
		print(searchresponse)
	}
}

func decodeMessage(message *kafka.Message) (request *constants.SearchCommunityRequest, s *status.Status) {
	if err := json.Unmarshal(message.Value, request); err != nil {
		return nil, status.GenerateBadRequest("poorly formatted data")
	}
	s = checkRequestForNecessaryData(request)
	return request, s
}

func checkRequestForNecessaryData(request *constants.SearchCommunityRequest) (s *status.Status) {
	if request.Name == "" && request.Tag == "" {
		return status.GenerateBadRequest("name and tag both cannot be empty")
	}
	if request.Pagesize == 0 {
		return status.GenerateBadRequest("page size cannot be null")
	}
	return status.OkStatus
}
