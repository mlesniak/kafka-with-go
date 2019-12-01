// Produces new messages by submitting them to the topic. Note that currently, random messages of a fixed length
// are sent.
package main

import (
	"encoding/hex"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"math"
	"math/rand"
)

func produce(broker string, number int, length int, topic string, group string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	for i := 1; i <= number; i++ {
		bytes := hex.EncodeToString(newRandomBytes(length))
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(bytes),
		}, nil)
		if err != nil {
			panic(err)
		}

		e := <-p.Events()
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			if i%*tick == 0 {
				log.Printf("Sending message %d/%d\n", i, number)
			}
		}
	}

	p.Flush(math.MaxInt64)
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
