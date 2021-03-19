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

func (s *pasteService) Get(ctx context.Context, ID string) error {
	return nil
}
