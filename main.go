package main

import "log"

func main() {
	log.Println("Starting")
	router, err := Inject()
	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}
	router.Run(":5000")
}
