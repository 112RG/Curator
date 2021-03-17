package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
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

	db, err := sql.Open("sqlite3", "./curator.db")
	if err != nil {
		panic(err.Error())
	}
	createTable(db)

	return db

}
func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS pastes (
		"idPaste" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"fId" TEXT,
		"content" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}
