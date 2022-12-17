package consumerlistener

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abhirajranjan/spaces/community/internal/consumer"
	"github.com/abhirajranjan/spaces/community/pkg/topics"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Listen() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false

		default:
			ev, err := consumer.Consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			go processEvent(ev)
		}
	}
}

func processEvent(message *kafka.Message) {
	switch *message.TopicPartition.Topic {
	case topics.GetCommunity:
		fmt.Println("Getting community")
		// GC.Handle(message)
	case topics.SearchCommunity:
		fmt.Println("Searching community")
		// SC.Handle(message)
	case topics.NewCommunity:
		fmt.Println("New community")
		// NC.Handle(message)
	case topics.NewSpace:
		fmt.Println("New Space")
		// NS.Handle(message)
	case topics.UpdateMetaData:
		fmt.Println("Updating metadata")
		// UM.Handle(message)
	}
}
