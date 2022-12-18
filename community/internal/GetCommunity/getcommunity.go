package getcommunity

import (
	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/internal/response"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"encoding/json"
)

func unmarshal(b []byte, out interface{}) error {
	err := json.Unmarshal(b, out)
	if err != nil {
		logger.Logger.Warn("error unmarshling data" + string(b))
	}
	return err
}

func Handle(message *kafka.Message) {
	var data constants.Community

	if unmarshal(message.Value, &data) != nil {
		return
	}

	if err := response.Send(message.Headers,
		generateResponse(db.Get(data))); err != nil {
		logger.Logger.Warn("error sending responses")
	}
}

func generateResponse(d primitive.D) string {
	return ""
}
