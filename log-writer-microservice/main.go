package main

import (
	"log-writer/kafka"
)

const (
	kafkaConn = "192.168.52.24:9092"
	topic     = "trendyol-example"
)

func main() {
	// kafka.CreateTopic(kafkaConn, topic)
	kafka.Producer(topic, kafkaConn)

}
