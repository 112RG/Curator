package model

import (
	"context"
	"database/sql"
	"time"
)

// User ..
type Paste struct {
	Id          string
	OwnerId     sql.NullString
	Expiry      sql.NullTime
	Title       sql.NullString
	TimeCreated time.Time
	Content     string
}

type PasteService interface {
	Get(ctx context.Context, Id string) (Paste, error)
	Delete(ctx context.Context, Id string) error
	Create(ctx context.Context, paste Paste) error
	GetOwnerPastes(ctx context.Context, Owner string) ([]*Paste, error)
}

type PasteRepository interface {
	FindByID(ctx context.Context, Id string) (Paste, error)
	DeleteByID(ctx context.Context, Id string) error
	CreatePaste(ctx context.Context, paste Paste) error
	FindByOwner(ctx context.Context, Owner string) ([]*Paste, error)
}
