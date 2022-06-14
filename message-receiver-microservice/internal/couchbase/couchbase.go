package couchbase

import (
	"fmt"
	"github.com/couchbase/gocb"
	"encoding/json"
)

var bucket *gocb.Bucket

type KafkaMessages struct{
    ID string `json:"id"`
	Message string `json:"message"`
}



func connectDB(){
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		fmt.Println("ERRROR CONNECTING TO CLUSTER:", err)
	}

	// Open Bucket
	bucket, err = cluster.OpenBucket("trendyol-example", "")
	if err != nil {
		fmt.Println("ERRROR OPENING BUCKET:", err)
	}
}
 
func CreateDocument(documentId string, message *KafkaMessages) {
	connectDB()
	fmt.Println("Upserting a full document...")
	_, error := bucket.Upsert(documentId, message, 0)
	if error != nil {
	 fmt.Println(error.Error())
	 return
	}
	getDocument(documentId)
}

 
func getDocument(documentId string) {
	fmt.Println("Getting the full document by id...")
	var message KafkaMessages
	_, error := bucket.Get(documentId, &message)
	if error != nil {
	 fmt.Println(error.Error())
	 return
	}
	jsonPerson, _ := json.Marshal(&message)
	fmt.Println(string(jsonPerson))
   }
	
    

	
