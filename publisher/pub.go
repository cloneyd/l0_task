package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	clusterID := "test-cluster"
	clientID := "test-pub"
	subj := "foo"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		log.Fatalf("Can't connect to nats-streaming: %s\n", err)
	}
	log.Printf("Connected to clusterID: [%s] clientID: [%s]\n", clusterID, clientID)

	byteValue, err := os.ReadFile("/json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %s\n", err)
	}

	err = sc.Publish(subj, byteValue) // does not return until an ack has been received from NATS Streaming
	if err != nil {
		log.Printf("Publish error: %s\n", err)
	} else {
		log.Println("Message has been sent")
	}

	// Close connection
	sc.Close()
}
