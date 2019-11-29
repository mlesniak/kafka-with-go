package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	producer = flag.Bool("produce", false, "Start producer")
	consumer = flag.Bool("consume", false, "Start consumer")
	group    = flag.String("group", "", "Set group for consumer")
	topic    = flag.String("topic", "", "Set topic")
	offset   = flag.String("offset", "end", "One of [beginning, end]")
	broker   = flag.String("broker", "localhost:9092", "Address of broker")
	number   = flag.Int("number", -1, "Number of messages to produce")
	length   = flag.Int("length", -1, "Length of a single message")
	tick     = flag.Int("tick", 1000, "Produce a log message every <tick> messages")
)

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

	if *consumer && *group == "" {
		fmt.Println("Group missing")
		flag.Usage()
		os.Exit(1)
	}

	if *number == -1 {
		fmt.Println("Number missing")
		flag.Usage()
		os.Exit(1)
	}

	if *producer && *length == -1 {
		fmt.Println("Length missing")
		flag.Usage()
		os.Exit(1)
	}
}
