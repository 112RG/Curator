package main

import (
	"github.com/112RG/Curator/injection"
	"log"
)

func main() {
	log.Println("Starting")
	router, err := injection.Inject()
	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}
	router.Run(":5000")
}
