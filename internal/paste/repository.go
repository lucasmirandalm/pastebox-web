package paste

import (
	"context"
	"database/sql"
)

type PasteRepository struct {
	db *sql.DB
}

func NewPasteRepository(db *sql.DB) *PasteRepository {
	return &PasteRepository{
		db: db,
	}
}

func (pr *PasteRepository) CountByUserID(ctx context.Context, userID int64) (int, error) {
	var total int

	err := pr.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM pastes WHERE user_id = $1`, userID).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
