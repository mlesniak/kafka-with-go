// Create a new topic.
package main

import (
	"context"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

// Replication factor is always 1 for now.
func createTopic(broker string, topic string, partitions int) {
	client, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	spec := kafka.TopicSpecification{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: 1,
		ReplicaAssignment: nil,
		Config:            nil,
	}
	log.Printf("Creating topic %s with %d partitions\n", topic, partitions)
	_, err = client.CreateTopics(ctx, []kafka.TopicSpecification{spec}, nil)
	if err != nil {
		panic(err)
	}
}
