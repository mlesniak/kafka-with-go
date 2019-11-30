package main

import (
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

	for counter := 1; counter <= number; counter++ {
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
