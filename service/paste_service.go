package service

import (
	"context"
	"errors"

	"curator/model"
)

type pasteService struct {
	PasteRepository model.PasteRepository
}
type USConfig struct {
	PasteRepository model.PasteRepository
}

func NewPasteService(c *USConfig) model.PasteService {
	return &pasteService{
		PasteRepository: c.PasteRepository,
	}
}

func (s *pasteService) Get(ctx context.Context, Id string) (model.Paste, error) {
	p, err := s.PasteRepository.FindByID(ctx, Id)
	return p, err
}

func (s *pasteService) Create(ctx context.Context, paste model.Paste) error {
	if len(paste.Content) > 1 {
		err := s.PasteRepository.CreatePaste(ctx, paste)
		return err
	} else {
		return errors.New("failed to create paste, content empty")
	}
}

func (s *pasteService) Delete(ctx context.Context, Id string) error {
	err := s.PasteRepository.DeleteByID(ctx, Id)
	return err
}

func (s *pasteService) GetOwnerPastes(ctx context.Context, Owner string) ([]*model.Paste, error) {
	pastes, err := s.PasteRepository.FindByOwner(ctx, Owner)
	return pastes, err
}
