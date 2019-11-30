package main

import (
	"encoding/hex"
	"golang.org/x/tools/container/intsets"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"math/rand"
)

func produce(broker string, number int, length int, topic string, group string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Fire and forget for now.
	for i := 1; i <= number; i++ {
		if i%*tick == 0 {
			// Wait until everything has been sent.
			p.Flush(intsets.MaxInt)
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
