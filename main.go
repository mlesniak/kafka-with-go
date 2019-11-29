package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func main() {
	//log.Println("Starting")
	//go listen("client-1")
	//go listen("client-2")
	//
	//time.Sleep(1000000 * 1000 * 10)
	//
	//go listen("client-3 / new")
	//
	//for {
	//	time.Sleep(1000 * 1000 * 3600)
	//}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "topic4"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	left := p.Flush(15 * 1000)
	log.Println("Left events:", left)
}

func listen(name string) {
	log.Println("Starting consumer", name)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": ":9092",
		"group.id":          "group",
		//"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{"topic4"}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("%s // Message on %s: %s\n", name, msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
}
