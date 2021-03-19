package service

import (
	"context"

	"github.com/112RG/Curator/model"
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
	err := s.PasteRepository.CreatePaste(ctx, paste)
	return err
}
func (s *pasteService) Delete(ctx context.Context, Id string) error {
	err := s.PasteRepository.DeleteByID(ctx, Id)
	return err
}
