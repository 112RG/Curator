package models

// User ..
type Paste struct {
	ID      string
	Content string
}

// UserRepository ..
type PasteRepository interface {
	FindByID(ID string) (*Paste, error)
	CreatePaste(paste *Paste) error
}
