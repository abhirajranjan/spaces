package consumerlistener

import (
	"encoding/base64"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abhirajranjan/spaces/community/internal/consumer"
	GC "github.com/abhirajranjan/spaces/community/internal/eventHandler/GetCommunity"
	NC "github.com/abhirajranjan/spaces/community/internal/eventHandler/NewCommunity"
	SC "github.com/abhirajranjan/spaces/community/internal/eventHandler/SearchCommunity"
	UC "github.com/abhirajranjan/spaces/community/internal/eventHandler/UpdateCommunity"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
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
	switch base64.RawStdEncoding.EncodeToString(message.Key) {
	case constants.GetCommunity:
		GC.Handle(message)
	case constants.SearchCommunity:
		SC.Handle(message)
	case constants.NewCommunity:
		NC.Handle(message)
	case constants.UpdateCommunity:
		UC.Handle(message)
	}
}
