package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"math/rand"
	"os"
)

// TODO Create new topics option

var (
	producer = flag.Bool("produce", false, "Start producer")
	consumer = flag.Bool("consume", false, "Start consumer")
	group    = flag.String("group", "", "Set group for consumer")
	topic    = flag.String("topic", "", "Set topic")
	broker   = flag.String("broker", "localhost:9092", "Address of broker")
	number   = flag.Int("number", -1, "Number of messages to produce")
	length   = flag.Int("length", -1, "Length of a single message")
)

func main() {
	initFlags()

	if *producer {
		log.Println("Starting producer")
		produce(*broker, *number, *length, *topic, *group)
		log.Println("Finished producer")
		return
	}

	if *consumer {
		log.Println("Consumer")
		return
	}
}

func initFlags() {
	flag.Parse()

	if !*producer && !*consumer {
		fmt.Println("Choose operation (-produce or -consume)")
		flag.Usage()
		os.Exit(1)
	}

	if broker == nil {
		fmt.Println("Broker missing")
		flag.Usage()
		os.Exit(1)
	}

	if topic == nil {
		fmt.Println("Topic missing")
		flag.Usage()
		os.Exit(1)
	}

	if group == nil {
		fmt.Println("Group missing")
		flag.Usage()
		os.Exit(1)
	}

	if *number == -1 {
		fmt.Println("Number missing")
		flag.Usage()
		os.Exit(1)
	}

	if *length == -1 {
		fmt.Println("Length missing")
		flag.Usage()
		os.Exit(1)
	}
}

func produce(broker string, number int, length int, topic string, group string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Fire and forget for now.
	for i := 1; i <= number; i++ {
		if i%(number/10) == 0 {
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

// newRandomBytes returns a new random array with the given length in bytes.
func newRandomBytes(length int) []byte {
	bs := make([]byte, length)
	_, err := rand.Read(bs)
	if err != nil {
		panic(err)
	}
	return bs
}
