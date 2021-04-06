package db

import (
	"database/sql"

	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func ConnectDB() *sql.DB {
	var err error
	if _, err := os.Stat("./curator.db"); os.IsNotExist(err) {
		log.Info().Msg("Creating curator.db...")

		file, err := os.Create("./curator.db") // Create SQLite file
		if err != nil {
			log.Error().Err(err)
		}
		file.Close()
		db, err := sql.Open("sqlite3", "./curator.db")
		if err != nil {
			panic(err.Error())
		}
		createPasteTable(db)
		createUserTable(db)
		createAlbumTable(db)
		log.Info().Msg("Created curator.db")

	}
	db, err := sql.Open("sqlite3", "./curator.db")
	if err != nil {
		panic(err.Error())
	}
	return db

}
func createPasteTable(db *sql.DB) {
	createPasteTable := `CREATE TABLE IF NOT EXISTS pastes (
		"paste_id" varchar(6) NOT NULL PRIMARY KEY,
		"album_id" varchar(10),
		"owner_id" varchar(36),
		"expiry" TIMESTAMP,
		"title" TEXT,
		"time_created" TIMESTAMP,
		"content" TEXT NOT NULL
	  );`

	createPasteTableIndex := `CREATE UNIQUE INDEX IF NOT EXISTS pastes_owner_id ON pastes (owner_id);`

	create, err := db.Prepare(createPasteTable)
	if err != nil {
		log.Error().Err(err)
	}
	create.Exec()

	index, err := db.Prepare(createPasteTableIndex)
	if err != nil {
		log.Error().Err(err)
	}
	index.Exec()
}

func createAlbumTable(db *sql.DB) {
	createAlbumTable := `CREATE TABLE IF NOT EXISTS albums (
		"album_id" varchar(10) NOT NULL PRIMARY KEY,
		"owner_id" varchar(36) NOT NULL,
		"title" TEXT NOT NULL
	);`

	createAlbumTableIndex := `CREATE UNIQUE INDEX IF NOT EXISTS album_owner_id ON albums (owner_id);`
	create, err := db.Prepare(createAlbumTable)
	if err != nil {
		log.Error().Err(err)
	}
	create.Exec()

	index, err := db.Prepare(createAlbumTableIndex)
	if err != nil {
		log.Error().Err(err)
	}
	index.Exec()
}
func createUserTable(db *sql.DB) {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		"user_id" varchar(6) NOT NULL PRIMARY KEY,
		"username" varchar(20) NOT NULL,
		"password" TEXT NOT NULL
	  );`
	create, err := db.Prepare(createUserTable)
	if err != nil {
		log.Error().Err(err)
	}
	create.Exec()

}
