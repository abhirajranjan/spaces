package updatespaces

import (
	"encoding/json"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(message *kafka.Message) {
	request, s := decodeMessage(message)

	switch s.Value {
	case status.BadRequestErrCode:
		// TODO: handle bad request
	case status.OkCode:
		res, s := db.UpdateSpaces(request)
		switch s.Value {
		case status.BadRequestErrCode:
			// TODO
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.UpdateSpaceRequest, s *status.Status) {
	if err := json.Unmarshal(message.Value, &request); err != nil {
		return nil, status.GenerateBadRequest("poorly formatted data")
	}

	s = checkRequestForNecessaryData(request)
	return request, s
}

func checkRequestForNecessaryData(request *constants.UpdateSpaceRequest) *status.Status {
	if request.Id == 0 {
		return status.GenerateBadRequest("id not specified")
	}
	return status.Ok
}
