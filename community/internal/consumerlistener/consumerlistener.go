package consumerlistener

import (
	"encoding/base64"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	GC "github.com/abhirajranjan/spaces/community/internal/GetCommunity"
	NC "github.com/abhirajranjan/spaces/community/internal/NewCommunity"
	NS "github.com/abhirajranjan/spaces/community/internal/NewSpace"
	SC "github.com/abhirajranjan/spaces/community/internal/SearchCommunity"
	UM "github.com/abhirajranjan/spaces/community/internal/UpdateMetadata"
	"github.com/abhirajranjan/spaces/community/internal/consumer"
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
	case constants.NewSpace:
		NS.Handle(message)
	case constants.UpdateMetaData:
		UM.Handle(message)
	}
}
