package updatemetadata

import (
	"encoding/json"
	"fmt"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// TODO: add validation of user to modify data
func Handle(message *kafka.Message) {
	request, _status := decodeMessage(message)
	switch _status.Value {
	case status.BadRequestErr:
		// TODO: handle err db
	case status.Ok:
		res, _status := db.UpdateCommunity(request)
		switch _status.Value {
		case status.InternalServerErr:
			// TODO: handle internal server error
		case status.Ok:
			// TODO: return updated data
			fmt.Println(res)
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.UpdatedCommunityRequest, s *status.Status) {
	if err := json.Unmarshal(message.Value, &request); err != nil {
		return nil, status.GenerateBadRequest("poorly formatted data")
	}
	return request, status.OkStatus
}
