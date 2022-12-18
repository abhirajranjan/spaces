package main

import (
	"github.com/abhirajranjan/spaces/community/config"
	"github.com/abhirajranjan/spaces/community/internal/consumer"
	"github.com/abhirajranjan/spaces/community/internal/consumerlistener"
	"github.com/abhirajranjan/spaces/community/internal/producer"
	"github.com/abhirajranjan/spaces/community/internal/producerlogger"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
)

func main() {
	producer.Initialize(config.KafkaProducerConfig())
	defer producer.Producer.Close()

	logger.InitializeLogger()
	go producerlogger.Producerlogger()

	consumer.Initialize(config.KafkaConsumerConfig())
	defer consumer.Consumer.Close()
	consumerlistener.Listen()
}
