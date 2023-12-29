package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost:9092",
		"broker.address.family": "v4",
		"client.id":             "foo",
		"acks":                  "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	deliveryChan := make(chan kafka.Event, 10000)
	topic := "quotes"

	for {
		quote := fetchQuote()
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          quote},
			deliveryChan,
		)
		if err != nil {
			fmt.Printf("Producer error: %v", err)
		}
		e := <-deliveryChan

		fmt.Printf("%+v\n", e.String())
		time.Sleep(2 * time.Second)
	}
}

func fetchQuote() []byte {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	return body
}
