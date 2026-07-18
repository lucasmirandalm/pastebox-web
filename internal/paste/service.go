package paste

import "context"

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
