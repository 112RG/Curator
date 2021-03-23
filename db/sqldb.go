package db

import (
	"database/sql"

	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func ConnectDB() *sql.DB {
	var err error
	log.Info().Msg("Starting DB")

	if _, err := os.Stat("./curator.db"); os.IsNotExist(err) {
		log.Info().Msg("Creating sqlite-database.db...")

		file, err := os.Create("./curator.db") // Create SQLite file
		if err != nil {
			log.Error().Err(err)
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", "./curator.db")
	if err != nil {
		panic(err.Error())
	}
	createTable(db)

	return db

}
func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS pastes (
		"_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Id" TEXT,
		"Expiry" INT,
		"Title" TEXT,
		"TimeCreated" TIMESTAMP,
		"CreatedIp" TEXT,
		"Owner" STRING,
		"Content" TEXT
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Error().Err(err)
	}
	statement.Exec() // Execute SQL Statements
}
