package main

import (
	"os"

	"github.com/abhirajranjan/spaces/community/config"
	"github.com/abhirajranjan/spaces/community/internal/consumer"
	"github.com/abhirajranjan/spaces/community/internal/consumerlistener"
	"github.com/abhirajranjan/spaces/community/internal/producer"
	"github.com/abhirajranjan/spaces/community/internal/producerlogger"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	var conf kafka.ConfigMap
	if len(os.Args) == 1 {
		conf = config.Load()
	} else {
		conf = config.Load(os.Args[1])
	}

	producer.Initialize(conf)
	defer producer.Producer.Close()
	logger.InitializeLogger()
	go producerlogger.Producerlogger()

	consumer.Initialize(conf)
	defer consumer.Consumer.Close()
	consumerlistener.Listen()
}
