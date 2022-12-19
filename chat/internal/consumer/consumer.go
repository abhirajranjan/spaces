package consumer

import (
	"log"
	"os"

	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Consumer *kafka.Consumer

func Initialize(c kafka.ConfigMap) {
	var err error

	var conf kafka.ConfigMap = c
	conf["group.id"] = constants.ConsumeGroupID
	conf["auto.offset.reset"] = "earliest"
	Consumer, err = kafka.NewConsumer(&conf)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
		os.Exit(1)
	}
	err = Consumer.SubscribeTopics([]string{constants.Self}, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
