package main

import (
	"fmt"
	kafka "github.com/Shopify/sarama"
)

func main() {
	fmt.Println("Hello, Kafka")

	config := kafka.NewConfig()
	client, err := kafka.NewClient([]string{":9092"}, config)
	if err != nil {
		panic(err)
	}

	producer(client)
}

func producer(client kafka.Client) {
	producer, _ := kafka.NewAsyncProducerFromClient(client)
	m := kafka.ProducerMessage{
		Topic: "topic",
		Value: kafka.StringEncoder("Hello, + ${time.Now()}"),
	}
	producer.Input() <- &m
	
	producer.Close()
}