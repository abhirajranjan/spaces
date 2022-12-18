package response

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO: implement response
func GenerateJson(doc bson.D) []byte
func Send(k []kafka.Header, s string) error
