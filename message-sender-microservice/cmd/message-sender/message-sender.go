package main

import (
	"message-sender-microservice/internal/kafka"
)

const (
	kafkaConn = "192.168.1.8:9092"
	topic     = "trendyol-example"
)

func main(){
	kafka.Producer(topic, kafkaConn)

}