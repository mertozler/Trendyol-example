package main

import (
	"fmt"
	"log"

	"github.com/couchbase/gocb/v2"
)

type Person struct {
	ID string `json:"id, omitempty"`
	Firstname string `json:"firstname, omitempty"`
	Lastname string `json:"lastname, omitempty"`
	Social []SocialMedia `json:"socialmedia, omitempty"`
}

type SocialMedia struct{
	Title string `json:"title, omitempty"`
	Link string `json:"link, omitempty"`
}


func main(){

	bucketName := "trendyol-example"
	username := "Administrator"
	password := "somepassword"

	// Initialize the Connection
	cluster, err := gocb.Connect("couchbase://192.168.1.172:11210", gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	bucket := cluster.Bucket(bucketName)

	

	col := bucket.Scope("trendyol-example").Collection("kafka-messages")
	type myDoc struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	document := myDoc{Foo: "bar", Bar: "foo"}
	result, err := col.Insert("document-key", &document, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)


	

}