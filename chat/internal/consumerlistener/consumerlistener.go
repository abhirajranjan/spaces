package consumerlistener

import (
	"encoding/base64"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	CR "github.com/abhirajranjan/spaces/chat/internal/eventHandler/CreateRoom"
	DC "github.com/abhirajranjan/spaces/chat/internal/eventHandler/DeleteChat"
	DR "github.com/abhirajranjan/spaces/chat/internal/eventHandler/DeleteRoom"
	NC "github.com/abhirajranjan/spaces/chat/internal/eventHandler/NewChat"
	RC "github.com/abhirajranjan/spaces/chat/internal/eventHandler/ReadChat"
	RU "github.com/abhirajranjan/spaces/chat/internal/eventHandler/RegisterUser"

	"github.com/abhirajranjan/spaces/chat/internal/consumer"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
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
			switch *ev.TopicPartition.Topic {
			case constants.Self:
				go processEvent(ev)
			case "user":
				go processUser(ev)
			}
		}
	}
}

func processEvent(message *kafka.Message) {
	switch base64.RawStdEncoding.EncodeToString(message.Key) {
	case constants.Event_NewChat:
		NC.Handle(message)
	case constants.Event_DeleteChat:
		DC.Handle(message)
	case constants.Event_CreateRoom:
		CR.Handle(message)
	case constants.Event_DeleteRoom:
		DR.Handle(message)
	case constants.Event_ReadChat:
		RC.Handle(message)
	}
}

func processUser(message *kafka.Message) {
	switch base64.RawStdEncoding.EncodeToString(message.Key) {
	case constants.Event_NewUserCreated:
		RU.Handle(message)
	}
}
