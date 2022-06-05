package main

import (
	"log-writer/kafka"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "trendyol-example"
)

func main() {
	// kafka.CreateTopic(kafkaConn, topic)
	kafka.Producer(topic, kafkaConn)

}
