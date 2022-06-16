package kafka

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	//"log"
	"os"
)

func Producer(topic string, kafkaConn string) {
	producer, err := InitProducer(kafkaConn)
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	// read command line input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish without goroutene
		Publish(topic, msg, producer)

		// publish with go routene
		// go publish(msg, producer)
	}
}

func InitProducer(kafkaConn string) (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	//sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func Publish(topic string, message string, producer sarama.SyncProducer) error{
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
	return nil
}
