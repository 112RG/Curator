package repositories

import (
	"context"
	"database/sql"

	"github.com/112RG/Curator/model"
	"github.com/rs/zerolog/log"
)

type pasteRepository struct {
	DB *sql.DB
}

func NewPasteRepository(db *sql.DB) model.PasteRepository {
	return &pasteRepository{
		DB: db,
	}
}

func (r *pasteRepository) FindByID(ctx context.Context, Id string) (p model.Paste, err error) {
	var mid sql.NullInt64
	err = r.DB.QueryRowContext(ctx, "SELECT * FROM pastes WHERE Id=?", Id).Scan(&mid, &p.Id, &p.Expiry, &p.Title, &p.TimeCreated, &p.CreatedIp, &p.Owner, &p.Content)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to find paste ID: %s", Id)
		return p, err
	} else {
		log.Debug().Msgf("Got paste ID: %s", Id)
		return p, err
	}
}

func (r *pasteRepository) CreatePaste(ctx context.Context, paste model.Paste) error {
	statement, err := r.DB.PrepareContext(ctx, `INSERT INTO pastes(Id, Content, TimeCreated, CreatedIp) VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to prepare SQL statement for ID: %s CONTENT: %s", paste.Id, paste.Content)
		return err
	} else {
		_, err = statement.ExecContext(ctx, paste.Id, paste.Content, paste.TimeCreated, paste.CreatedIp)
		return err
	}
}
func (r *pasteRepository) DeleteByID(ctx context.Context, Id string) error {
	statement, err := r.DB.Prepare(`DELETE FROM pastes WHERE Id=?`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete paste ID: %s", Id)
		return err
	} else {
		log.Info().Msgf("Deleted paste ID: %s", Id)
		_, err = statement.Exec(Id)
		return err
	}
}
