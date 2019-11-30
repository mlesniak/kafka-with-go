package main

import (
	"log"
)

// TODO Create new topics option
// TODO Refactoring

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
