package main

import (
	"github.com/112RG/Curator/routes"
)

func main() {
	// Our server will live in the routes package
	router := routes.Build()
	router.Run(":5000")
}
