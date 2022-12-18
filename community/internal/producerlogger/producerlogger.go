package producerlogger

import (
	"fmt"

	"github.com/abhirajranjan/spaces/community/internal/producer"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producerlogger() {
	for e := range producer.Producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				logger.Logger.Sugar().Warnf("Failed to deliver message: %v\n", ev.TopicPartition.Topic)
			} else {
				// ? should we really log that ?
				go fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			}
		}
	}
}
