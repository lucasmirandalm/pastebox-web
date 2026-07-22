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

func (ps *PasteService) ListByUserID(ctx context.Context, userID int64, onlyFavorites bool) ([]Paste, error) {
	return ps.repository.ListByUserID(ctx, userID, onlyFavorites)
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

func (ps *PasteService) Create(
	ctx context.Context,
	userID int64,
	title string,
	content string,
) (Paste, error) {
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)

	if title == "" {
		return Paste{}, ErrPasteTitleRequired
	}

	if content == "" {
		return Paste{}, ErrPasteContentRequired
	}

	publicID, err := generatePublicID()
	if err != nil {
		return Paste{}, err
	}

	return ps.repository.Create(ctx, userID, title, content, publicID)
}
