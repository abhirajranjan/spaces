package logger

import (
	"fmt"

	"github.com/abhirajranjan/spaces/community/internal/producer"
	"github.com/abhirajranjan/spaces/community/pkg/topics"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

type customwriter struct {
	Producer *kafka.Producer
}

func InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	mywriter := customwriter{producer.Producer}
	writer := zapcore.AddSync(mywriter)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func (ls customwriter) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	err = ls.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topics.Log, Partition: kafka.PartitionAny},
		Key:            []byte(topics.SelfPrefix),
		Value:          []byte(p),
	}, nil)
	n = len(p)
	return n, err
}
