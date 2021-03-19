package main

import (
	"github.com/112RG/Curator/db"
	"github.com/112RG/Curator/routes"
)

func main() {
	// Build reqs for router
	db := db.ConnectDB()
	router := routes.Build(db)
	router.Run(":5000")
}
