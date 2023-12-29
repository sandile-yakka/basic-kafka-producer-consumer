package main

import (
	"dummy/types"
	"fmt"
	"os"

	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	fmt.Println("Starting Consumer...")
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost:9092",
		"broker.address.family": "v4",
		"group.id":              "foo",
		"auto.offset.reset":     "smallest"})
	if err != nil {
		fmt.Printf("Init error: %+v", err)
	}
	err = consumer.Subscribe("quotes", nil)
	if err != nil {
		fmt.Printf("Subscribe error: %+v", err)
	}

	for {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			quote := types.Quote{}
			json.Unmarshal([]byte(e.Value), &quote)
			fmt.Printf("Message: %+v \n", quote)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "----- %% Error: %v\n", e)
		default:
			fmt.Printf("Ignored %+v\n", e)
		}
	}
}
