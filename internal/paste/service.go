package paste

import (
	"context"
	"strings"
)

type PasteService struct {
	repository *PasteRepository
}

func NewPasteService(repository *PasteRepository) *PasteService {
	return &PasteService{
		repository: repository,
	}
}

func (ps *PasteService) CountByUserID(ctx context.Context, userID int64) (int, error) {
	return ps.repository.CountByUserID(ctx, userID)
}

func (ps *PasteService) ListByUserID(ctx context.Context, userID int64) ([]Paste, error) {
	return ps.repository.ListByUserID(ctx, userID)
}

func (ps *PasteService) FindByID(ctx context.Context, userID, pasteID int64) (Paste, error) {
	return ps.repository.FindByID(ctx, userID, pasteID)
}

func (ps *PasteService) Update(
	ctx context.Context,
	userID int64,
	pasteID int64,
	title string,
	content string,
	isFavorite bool,
) (Paste, error) {
	if strings.TrimSpace(title) == "" {
		return Paste{}, ErrPasteTitleRequired
	}

	if strings.TrimSpace(content) == "" {
		return Paste{}, ErrPasteContentRequired
	}

	return ps.repository.Update(ctx, userID, pasteID, title, content, isFavorite)
}
