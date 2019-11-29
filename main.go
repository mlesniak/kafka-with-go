package main

import (
	"context"
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
	consumer(client)
	//consumerGroup(client)
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
	err = group.Consume(ctx, []string{"topic"}, groupHandler{})
	if err != nil {
		panic(err)
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

type groupHandler struct{}

func (groupHandler) Setup(_ kafka.ConsumerGroupSession) error   { return nil }
func (groupHandler) Cleanup(_ kafka.ConsumerGroupSession) error { return nil }
func (h groupHandler) ConsumeClaim(sess kafka.ConsumerGroupSession, claim kafka.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}
