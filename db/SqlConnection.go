package connection

import (
	"database/sql"
	"log"
	"os"
)

var SQLConn *sql.DB

func setupDatabase() {
	var err error
	if _, err := os.Stat("./curator.db"); os.IsNotExist(err) {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("./curator.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")
	}

	SQLConn, err = sql.Open("sqlite3", "./curator.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer SQLConn.Close()

}
