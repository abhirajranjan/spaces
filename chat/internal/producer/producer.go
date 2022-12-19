package producer

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer *kafka.Producer

func Initialize(conf kafka.ConfigMap) {
	var err error
	Producer, err = kafka.NewProducer(&conf)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
		os.Exit(1)
	}
}
