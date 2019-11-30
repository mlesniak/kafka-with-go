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
}
