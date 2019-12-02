// Handle command line parsing.
package main

import (
	"flag"
	"log"
	"os"
)

var (
	producer   = flag.Bool("produce", false, "Start producer")
	consumer   = flag.Bool("consume", false, "Start consumer")
	create     = flag.Bool("create", false, "Create a new topic")
	list       = flag.Bool("list", false, "List available topics")
	group      = flag.String("group", "", "Set group for consumer")
	topic      = flag.String("topic", "", "Set topic")
	broker     = flag.String("broker", getBroker(), "Address of broker")
	number     = flag.Int("number", -1, "Number of messages to produce")
	length     = flag.Int("length", -1, "Length of a single message")
	tick       = flag.Int("tick", 1000, "Produce a log message every <tick> messages")
	partitions = flag.Int("partitions", 1, "Number of partitions")
)

// getBroker returns the broker to use. By default it uses localhost for development systems, but if an environment
// variable BROKER exists, it is used instead.
func getBroker() string {
	broker, ok := os.LookupEnv("BROKER")
	if ok {
		log.Printf("Using broker from environment variable BROKER: %s\n", broker)
		return broker
	}
	return "localhost:9092"
}

// initFlags checks correct set of command line flags for the different operations.
func initFlags() {
	flag.Parse()

	type flagRule struct {
		condition bool
		message   string
	}

	rules := []flagRule{
		{!*producer && !*consumer && !*create && !*list, "Choose operation (-produce or -consume or -create)"},
		{broker == nil, "Broker missing"},
		{(*producer || *consumer || *create) && topic == nil, "Topic missing"},
		{*consumer && *group == "", "Group missing"},
		{*producer && *number == -1, "Number missing"},
		{*producer && *length == -1, "Length missing"},
	}

	for _, rule := range rules {
		if rule.condition {
			log.Println(rule.message + "\n")
			flag.Usage()
			os.Exit(1)
		}
	}
}
