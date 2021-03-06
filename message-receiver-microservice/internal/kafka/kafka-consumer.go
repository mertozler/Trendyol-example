package kafka

import (
	"fmt"
	"log"
	"message-receiver-microservice/internal/couchbase"
	"os"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "trendyol-example"
)

func Consumer() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// init consumer
	cg, err := initConsumer()
	if err != nil {
		fmt.Println("Error consumer goup: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	consume(cg)
}

func initConsumer() (*consumergroup.ConsumerGroup, error) {
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup("zgroup", []string{topic}, []string{"localhost:2181"}, config)
	if err != nil {
		return nil, err
	}

	return cg, err
}

func consume(cg *consumergroup.ConsumerGroup) {
	counter := 1
	for {
		select {
		case msg := <-cg.Messages():
			// messages coming through chanel
			// only take messages from subscribed topic
			if msg.Topic != topic {
				continue
			}
			newMessage := couchbase.KafkaMessages{
				ID: strconv.Itoa(counter),
				Message: string(msg.Value),
			}
			
			couchbase.CreateDocument(strconv.Itoa(counter), &newMessage)
			
			counter = counter + 1
			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
