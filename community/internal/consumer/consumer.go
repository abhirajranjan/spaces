package consumer

import (
	"log"
	"os"

	"github.com/abhirajranjan/spaces/community/pkg/topics"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Consumer *kafka.Consumer

func Initialize(c kafka.ConfigMap) {
	var err error

	var conf kafka.ConfigMap = c
	conf["group.id"] = "kafka-go-getting-started"
	conf["auto.offset.reset"] = "earliest"
	Consumer, err = kafka.NewConsumer(&conf)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
		os.Exit(1)
	}
	err = Consumer.SubscribeTopics([]string{"^" + topics.SelfPrefix + ".*"}, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
