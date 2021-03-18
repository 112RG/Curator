package repositories

import (
	"database/sql"
	"log"

	"github.com/112RG/Curator/models"
)

// UserRepo implements models.UserRepository
type PasteRepo struct {
	db *sql.DB
}

// NewUserRepo ..
func NewPasteRepo(db *sql.DB) *PasteRepo {
	return &PasteRepo{
		db: db,
	}
}

// FindByID ..
func (r *PasteRepo) FindByID(ID string) (u models.Paste, err error) {
	var p models.Paste
	var mid sql.NullInt32
	err = r.db.QueryRow("SELECT * FROM pastes WHERE fId=?", ID).Scan(&mid, &p.ID, &p.Content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return p, err
}

// Save ..
func (r *PasteRepo) CreatePaste(paste *models.Paste) error {
	statement, err := r.db.Prepare(`INSERT INTO pastes(fId, content) VALUES (?, ?)`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(paste.ID, paste.Content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err
}
