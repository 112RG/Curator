package repositories

import (
	"database/sql"

	"github.com/112RG/Curator/models"
	"github.com/rs/zerolog/log"
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
	return p, err
}

// Save ..
func (r *PasteRepo) CreatePaste(paste *models.Paste) error {
	statement, err := r.db.Prepare(`INSERT INTO pastes(fId, content) VALUES (?, ?)`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to prepare SQL statement for ID: %s CONTENT: %s", paste.ID, paste.Content)
		return err
	} else {
		_, err = statement.Exec(paste.ID, paste.Content)
	}
	return err
}
