package main

import (
	"context"
	"log"

	//"context"
	"fmt"
	kafka "github.com/Shopify/sarama"
)

func main() {
	config := kafka.NewConfig()
	config.Version = kafka.MaxVersion
	client, err := kafka.NewClient([]string{":9092"}, config)
	if err != nil {
		panic(err)
	}

	// producer(client)
	//consumer(client)
	consumerGroup(client)
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

func consumerGroup(client kafka.Client) {
	v := client.Config().Version
	fmt.Println(v)

	group, err := kafka.NewConsumerGroupFromClient("readers", client)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Messages")
	ctx := context.Background()
	for {
		// Simulate two consumer.
		go func() {
			fmt.Println("Starting parallel consumer")
			group.Consume(ctx, []string{"topic4"}, groupHandler{"client 0"})
		}()
		group.Consume(ctx, []string{"topic4"}, groupHandler{"client 1"})
	}
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

	part, err := consumer.ConsumePartition("topic4", 0, kafka.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Messages")
	messages := part.Messages()
	for m := range messages {
		fmt.Println(string(m.Value))
	}
}

type groupHandler struct {
	Name string
}

func (g groupHandler) Setup(session kafka.ConsumerGroupSession) error {
	fmt.Println("Started", g.Name)
	fmt.Println("Claims:", session.Claims())
	return nil
}
func (groupHandler) Cleanup(_ kafka.ConsumerGroupSession) error { return nil }
func (h groupHandler) ConsumeClaim(sess kafka.ConsumerGroupSession, claim kafka.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		sess.MarkMessage(msg, "")
		if sess.Context().Err() != nil {
			log.Println("Context switch")
			return nil
		}
	}
	return nil
}
