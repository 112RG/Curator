package repositories

import (
	"database/sql"

	"github.com/112RG/Curator/model"
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
func (r *PasteRepo) FindByID(Id string) (u model.Paste, err error) {
	var p model.Paste
	var mid sql.NullInt64
	err = r.db.QueryRow("SELECT * FROM pastes WHERE Id=?", Id).Scan(&mid, &p.Id, &p.Expiry, &p.Title, &p.TimeCreated, &p.CreatedIp, &p.Owner, &p.Content)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to find paste ID: %s", Id)
		return p, err
	} else {
		log.Debug().Msgf("Got paste ID: %s", Id)
		return p, err
	}
}

// Save ..
/* func (r *PasteRepo) CreatePaste(paste *model.Paste) error {
	statement, err := r.db.Prepare(`INSERT INTO pastes(Id, Content, TimeCreated, CreatedIp) VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to prepare SQL statement for ID: %s CONTENT: %s", paste.Id, paste.Content)
		return err
	} else {
		_, err = statement.Exec(paste.Id, paste.Content, paste.TimeCreated, paste.CreatedIp)
		return err
	}
} */
/*
func (r *PasteRepo) DeletePasteByID(ID string) error {
	statement, err := r.db.Prepare(`DELETE FROM pastes WHERE fId=?`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete paste ID: %s", ID)
		return err
	} else {
		log.Info().Msgf("Deleted paste ID: %s", ID)
		_, err = statement.Exec(ID)
		return err
	}
}
*/
