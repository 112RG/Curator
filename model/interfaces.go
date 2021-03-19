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
	Get(ctx context.Context) error
	//Delete(ID string) error
	//CreatePaste(p *Paste) error
}

type PasteRepository interface {
	FindByID(ID string) (*Paste, error)
	//DeleteByID(ID string) error
	//Create(paste *Paste) error
}
