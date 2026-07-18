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

func (pr *PasteRepository) ListByUserID(ctx context.Context, userID int64) ([]Paste, error) {
	rows, err := pr.db.QueryContext(ctx, `
		SELECT id, user_id, title, content, is_favorite, public_id, created_at, updated_at
		FROM pastes
		WHERE user_id = $1
		ORDER BY updated_at DESC, id DESC
		LIMIT 10
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pastes []Paste

	for rows.Next() {
		var paste Paste

		err := rows.Scan(
			&paste.ID,
			&paste.UserID,
			&paste.Title,
			&paste.Content,
			&paste.IsFavorite,
			&paste.PublicID,
			&paste.CreatedAt,
			&paste.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		pastes = append(pastes, paste)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pastes, nil
}
