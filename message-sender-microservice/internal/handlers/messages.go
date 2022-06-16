package handlers

import (
	log "github.com/sirupsen/logrus"
	"message-sender-microservice/internal/kafka"
	"net/http"
	"message-sender-microservice/internal/models"
)

const (
	kafkaConn = "192.168.1.8:9092"
	topic     = "trendyol-example"
)

type Messages struct{
	l *log.Logger
}

func NewMessages(l *log.Logger) *Messages{
	return &Messages{l}
}

func (m *Messages) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	if(r.Method == http.MethodGet){
		m.l.Println("Healtcheck status is fine")
		rw.Write([]byte("Hello world"))
		return
	}

	if(r.Method == http.MethodPost){
		m.l.Println("Received a message")
		m.addMessage(rw, r)
		return
	}


	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (m *Messages) addMessage(rw http.ResponseWriter, r*http.Request){
	m.l.Println("sending message to kafka")
	producer, err := kafka.InitProducer(kafkaConn)
	if err != nil {
		m.l.Error("Error producer: ", err.Error())
		http.Error(rw, "An unknown error occurred while sending the message", http.StatusInternalServerError)
		return
	}
	message := &models.Message{}
	err = message.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
	}
	err = kafka.Publish(topic,message.Message,producer)
	if err != nil{
		m.l.Error("Error publishing message ", err.Error())
		http.Error(rw, "An unknown error occurred while sending the message", http.StatusInternalServerError)
		return
	}
	m.l.Println("message sent successfully")
}