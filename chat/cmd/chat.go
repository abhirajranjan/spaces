package main

import (
	"github.com/abhirajranjan/spaces/chat/config"
	"github.com/abhirajranjan/spaces/chat/internal/consumer"
	"github.com/abhirajranjan/spaces/chat/internal/consumerlistener"
	_ "github.com/abhirajranjan/spaces/chat/internal/db" // init database
	"github.com/abhirajranjan/spaces/chat/internal/producer"
	"github.com/abhirajranjan/spaces/chat/internal/producerlogger"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
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
