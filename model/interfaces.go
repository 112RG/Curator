package model

import (
	"context"
	"database/sql"
	"time"
)

// User ..
type Paste struct {
	Expiry      sql.NullInt64
	Title       sql.NullString
	TimeCreated time.Time
	CreatedIp   string
	Owner       sql.NullString
	Content     string
	Id          string
}

type PasteService interface {
	Get(ctx context.Context, ID string) (Paste, error)
	//Delete(ID string) error
	Create(ctx context.Context, paste Paste) error
}

type PasteRepository interface {
	FindByID(ctx context.Context, ID string) (Paste, error)
	DeleteByID(ID string) error
	CreatePaste(ctx context.Context, paste Paste) error
}
