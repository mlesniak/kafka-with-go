package main

import (
	"fmt"
	kafka "github.com/Shopify/sarama"
)

func main() {
	config := kafka.NewConfig()
	client, err := kafka.NewClient([]string{":9092"}, config)
	if err != nil {
		panic(err)
	}

	// producer(client)
	consumer(client)
}

func producer(client kafka.Client) {
	producer, _ := kafka.NewAsyncProducerFromClient(client)
	m := kafka.ProducerMessage{
		Topic: "topic",
		Value: kafka.StringEncoder("Hello, + ${time.Now()}"),
	}
	producer.Input() <- &m
	// TODO Read error channel to prevent blocking
	
	producer.Close()
}

func consumer(client kafka.Client) {
	consumer, err := kafka.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}

	topics, _ := consumer.Topics()
	fmt.Println("--- Topics")
	for _, value := range topics {
		fmt.Println(value)
	}

	part, err := consumer.ConsumePartition("topic", 0, kafka.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Messages")
	messages := part.Messages()
	for m := range messages {
		fmt.Println(string(m.Value))
	}
}