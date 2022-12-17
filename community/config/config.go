package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func readConfig(configFile string) kafka.ConfigMap {
	fmt.Println(configFile)

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

func Load(configFile ...string) kafka.ConfigMap {
	if len(configFile) != 0 {
		return readConfig(configFile[0])
	} else {
		return readConfig("kafka.properties")
	}
}
