package main

import (
	"encoding/hex"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"math/rand"
)

// TODO Create new topics option
// TODO Refactoring
// TODO multiple consumers

func main() {
	initFlags()

	if *producer {
		log.Println("Starting producer")
		produce(*broker, *number, *length, *topic, *group)
		log.Println("Finished producer")
		return
	}

	if *consumer {
		log.Println("Starting consumer")
		consume(*broker, *topic, *group, *number)
		log.Println("Finished consumer")
		return
	}
}

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

func produce(broker string, number int, length int, topic string, group string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Fire and forget for now.
	for i := 1; i <= number; i++ {
		if i%*tick == 0 {
			log.Printf("Sending message %d/%d\n", i, number)
		}
		bytes := hex.EncodeToString(newRandomBytes(length))
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(bytes),
		}, nil)
		if err != nil {
			panic(err)
		}
	}
}

// newRandomBytes returns a new random array with the given length in bytes.
func newRandomBytes(length int) []byte {
	bs := make([]byte, length)
	_, err := rand.Read(bs)
	if err != nil {
		panic(err)
	}
	return bs
}
