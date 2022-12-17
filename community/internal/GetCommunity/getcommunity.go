package getcommunity

import (
	"github.com/abhirajranjan/spaces/community/internal/db"
	"github.com/abhirajranjan/spaces/community/internal/response"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/abhirajranjan/spaces/community/pkg/topics"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"encoding/json"
)

func Handle(message *kafka.Message) {
	var jkey topics.Community
	err := json.Unmarshal(message.Key, &jkey)
	if err != nil {
		logger.Logger.Warn("error unmarshling data")
	}
	response.Send(message.Headers,
		response.GenerateJson(db.GetCommunity(jkey)),
	)
}
