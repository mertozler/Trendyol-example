package main

import (
	
	//"message-sender-microservice/internal/kafka"
	"message-sender-microservice/internal/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
	"context"
	log "github.com/sirupsen/logrus"
)



func main(){
	
	//kafka.Producer(topic, kafkaConn)
	l := log.New()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})

	mes := handlers.NewMessages(l)

	sm := http.NewServeMux()
	sm.Handle("/",mes)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	
	go func(){
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)


}