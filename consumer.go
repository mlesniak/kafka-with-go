// Consume submitted messages and list available topics (since both are consume-operations, we define them in the
// same file). Note that we discard the messages for now.
package main

import (
	"golang.org/x/tools/container/intsets"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func consume(broker string, topic string, group string, number int) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		// We want to read all stored messages if we have not started yet for a group.
		"auto.offset.reset": "beginning",
	})
	if err != nil {
		panic(err)
	}

	_ = c.SubscribeTopics([]string{topic}, nil)

	for counter := 1; number == -1 || counter <= number; counter++ {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			if counter%*tick == 0 {
				log.Printf("Received %d messages (%d new)\n", counter, *tick)
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	_ = c.Close()
	log.Printf("Consumer read %d entries\n", number)
}

func listTopics(broker string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "",
	})
	if err != nil {
		panic(err)
	}

	metadata, err := c.GetMetadata(nil, false, intsets.MaxInt)
	if err != nil {
		panic(err)
	}
	for topic, value := range metadata.Topics {
		log.Printf("%s /%d\n", topic, len(value.Partitions))
	}
}
