// Playground examples for using kafka with go. The general workflow is as follows: 1) create topic 2) check success
// by listing it 3) produce some data and 4) consume data.
package main

import (
	"log"
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
		log.Println("Starting consumer")
		consume(*broker, *topic, *group, *number)
		log.Println("Finished consumer")
		return
	}

	if *create {
		log.Println("Starting topic creation")
		createTopic(*broker, *topic, *partitions)
		log.Println("Finished topic creation")
		return
	}

	if *list {
		log.Println("Listing topics")
		listTopics(*broker)
		return
	}
}
