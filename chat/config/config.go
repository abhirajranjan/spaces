package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var dbURI string
var dbToken string
var kafkaProducerConfig kafka.ConfigMap
var kafkaConsumerConfig kafka.ConfigMap

var kafkaProdMutex = &sync.Mutex{}
var kafkaConsMutex = &sync.Mutex{}
var dbURIMutex = &sync.Mutex{}
var dbTokenMutex = &sync.Mutex{}

// getter functions
func DbURI() string {
	dbURIMutex.Lock()
	defer dbURIMutex.Unlock()
	return dbURI
}

func DbToken() string {
	dbTokenMutex.Lock()
	defer dbTokenMutex.Unlock()
	return dbToken
}

func KafkaProducerConfig() kafka.ConfigMap {
	kafkaProdMutex.Lock()
	defer kafkaProdMutex.Unlock()
	return kafkaProducerConfig
}

func KafkaConsumerConfig() kafka.ConfigMap {
	kafkaConsMutex.Lock()
	defer kafkaConsMutex.Unlock()
	return kafkaConsumerConfig
}

// log formatting function
func logErr(cause string, err error) {
	if cause != "" {
		if err != nil {
			log.Fatalf("%s:%v\n", cause, err)
		}
		log.Fatalf("%s\n", cause)
	}
	log.Fatal(err)
}

// load database from env
func loadDB() {
	clusterid, ok1 := os.LookupEnv(constants.ENV_astraClusterId)
	region, ok2 := os.LookupEnv(constants.ENV_astraRegion)
	token, ok3 := os.LookupEnv(constants.ENV_BearerToken)

	if ok1 && ok2 && ok3 {
		dbTokenMutex.Lock()
		dbURIMutex.Lock()
		defer dbURIMutex.Unlock()
		defer dbTokenMutex.Unlock()
		dbURI = fmt.Sprintf(constants.DB_URI, clusterid, region)
		dbToken = token
	} else {
		logErr(fmt.Sprintf("cannot locate enviroment variable %s or %s or %s\n",
			constants.ENV_BearerToken,
			constants.ENV_astraClusterId,
			constants.ENV_astraRegion,
		), nil)
	}
}

// load producer config file from env
func loadProducer() {
	if value, ok := os.LookupEnv(constants.ENV_producerConfig); ok {
		kafkaProdMutex.Lock()
		defer kafkaProdMutex.Unlock()
		kafkaProducerConfig = readConfig(value)
	} else {
		logErr(fmt.Sprintf("cannot locate enviroment varible %s\n", constants.ENV_producerConfig), nil)
	}
}

// load consumer config file from env
func loadConsumer() {
	if value, ok := os.LookupEnv(constants.ENV_consumerConfig); ok {
		kafkaConsMutex.Lock()
		defer kafkaConsMutex.Unlock()
		kafkaConsumerConfig = readConfig(value)
	} else {
		logErr(fmt.Sprintf("cannot locate enviroment varible %s\n", constants.ENV_consumerConfig), nil)
	}
}

// open config file and convert it into kafka.ConfigMap
func readConfig(configFile string) kafka.ConfigMap {
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %s", err)
		os.Exit(1)
	}
	return m
}

// load all configs
func init() {
	loadProducer()
	loadConsumer()
	loadDB()
}
