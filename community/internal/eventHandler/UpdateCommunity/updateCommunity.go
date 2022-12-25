package updatemetadata

import (
	"encoding/json"
	"fmt"

	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// TODO: add validation of user to modify data
func Handle(message *kafka.Message) {
	request, status := decodeMessage(message)
	switch status.Value {
	case constants.BadRequestErr:
		// TODO: handle err db
	case constants.Ok:
		res, status := db.UpdateCommunity(request)
		switch status.Value {
		case constants.InternalServerErr:
			// TODO: handle internal server error
		case constants.Ok:
			// TODO: return updated data
			fmt.Println(res)
		}
	}
}

func decodeMessage(message *kafka.Message) (request *constants.UpdatedCommunityRequest, status *constants.Status) {
	if err := json.Unmarshal(message.Value, &request); err != nil {
		return nil, constants.GenerateBadRequest("poorly formatted data")
	}
	return request, constants.Status_Ok
}
