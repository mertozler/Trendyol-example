package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"golang.org/x/exp/slices"
	"log"
)

func CreateTopic(kafkaConn string, topic string) {
	brokerAddrs := []string{kafkaConn}
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()
	cluster, err := sarama.NewConsumer(brokerAddrs, config)
	if err != nil {
		panic(err)
	}
	topics, _ := cluster.Topics()
	if slices.Contains(topics, topic) {
		fmt.Println("Topic is already existing")
	} else {
		err = admin.CreateTopic(topic, &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		}, false)
		if err != nil {
			log.Fatal("Error while creating topic: ", err.Error())

		}
	}
}
