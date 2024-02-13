package repositories

import (
	"context"
	"database/sql"

	"curator/model"

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

func (r *pasteRepository) FindByID(ctx context.Context, Id string) (model.Paste, error) {
	p := model.Paste{}
	err := r.DB.QueryRowContext(ctx, "SELECT * FROM pastes WHERE paste_id=?", Id).Scan(&p.Id, &p.AlbumId, &p.OwnerId, &p.Lang, &p.Expiry, &p.Title, &p.TimeCreated, &p.Content)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to find paste ID: %s", Id)
		return p, err
	} else {
		log.Debug().Msgf("Got paste ID: %s", Id)
		return p, nil
	}
}

func (r *pasteRepository) CreatePaste(ctx context.Context, paste model.Paste) error {
	log.Debug().Msg(paste.Title.String)
	log.Debug().Msg(paste.Expiry.Time.String())

	statement, err := r.DB.PrepareContext(ctx, `INSERT INTO pastes(paste_id, owner_id, album_id, lang, expiry, title, time_created, content) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to prepare SQL statement for ID: %s CONTENT: %s", paste.Id, paste.Content)
	} else {
		_, err = statement.ExecContext(ctx, paste.Id, NewNullString(paste.OwnerId.String), NewNullString(paste.AlbumId.String), paste.Lang, paste.Expiry.Time, NewNullString(paste.Title.String), paste.TimeCreated, paste.Content)
	}
	return err
}

func (r *pasteRepository) DeleteByID(ctx context.Context, Id string) error {
	statement, err := r.DB.Prepare(`DELETE FROM pastes WHERE paste_id=?`)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete paste ID: %s", Id)
		return err
	} else {
		log.Info().Msgf("Deleted paste ID: %s", Id)
		_, err = statement.Exec(Id)
		return err
	}
}

func (r *pasteRepository) FindByOwner(ctx context.Context, Owner string) (p []*model.Paste, err error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM pastes WHERE Owner=?", Owner)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to query db PARAM: %s", Owner)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		paste := new(model.Paste)
		if err := rows.Scan(&paste.Id, &paste.OwnerId.String, &paste.Expiry, &paste.Title, &paste.TimeCreated, &paste.Content); err != nil {
			log.Error().Err(err).Msgf("Failed to scan row: %s", Owner)
		} else {
			p = append(p, paste)
		}
	}
	return p, err
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
