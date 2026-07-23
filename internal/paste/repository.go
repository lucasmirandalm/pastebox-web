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

func (pr *PasteRepository) CountByUserID(ctx context.Context, userID int64, onlyFavorites bool, search string) (int, error) {
	var total int

	err := pr.db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM pastes
		WHERE user_id = $1
		  AND ($2 = false OR is_favorite = true)
		  AND (
		    $3 = ''
		    OR title ILIKE '%' || $3 || '%'
		    OR content ILIKE '%' || $3 || '%'
		  )
	`, userID, onlyFavorites, search).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (pr *PasteRepository) ListByUserID(ctx context.Context, userID int64, onlyFavorites bool, search string) ([]Paste, error) {
	rows, err := pr.db.QueryContext(ctx, `
		SELECT id, user_id, title, content, is_favorite, public_id, created_at, updated_at
		FROM pastes
		WHERE user_id = $1
		  AND ($2 = false OR is_favorite = true)
		  AND (
		    $3 = ''
		    OR title ILIKE '%' || $3 || '%'
		    OR content ILIKE '%' || $3 || '%'
		  )
		ORDER BY updated_at DESC, id DESC
		LIMIT 10
	`, userID, onlyFavorites, search)
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

func (pr *PasteRepository) FindByID(ctx context.Context, userID int64, pasteID int64) (Paste, error) {
	var paste Paste

	err := pr.db.QueryRowContext(ctx, `
		SELECT id, user_id, title, content, is_favorite, public_id, created_at, updated_at
		FROM pastes
		WHERE id = $1
		  AND user_id = $2
	`, pasteID, userID).Scan(
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
		if err == sql.ErrNoRows {
			return Paste{}, ErrPasteNotFound
		}

		return Paste{}, err
	}

	return paste, nil
}

func (pr *PasteRepository) Update(
	ctx context.Context,
	userID int64,
	pasteID int64,
	title string,
	content string,
	isFavorite bool,
) (Paste, error) {
	var paste Paste

	err := pr.db.QueryRowContext(ctx, `
		UPDATE pastes
		SET
			title = $3,
			content = $4,
			is_favorite = $5,
			updated_at = now()
		WHERE id = $1
		  AND user_id = $2
		RETURNING id, user_id, title, content, is_favorite, public_id, created_at, updated_at
	`, pasteID, userID, title, content, isFavorite).Scan(
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
		if err == sql.ErrNoRows {
			return Paste{}, ErrPasteNotFound
		}

		return Paste{}, err
	}

	return paste, nil
}

func (pr *PasteRepository) Create(
	ctx context.Context,
	userID int64,
	title string,
	content string,
	publicID string,
) (Paste, error) {
	var paste Paste

	err := pr.db.QueryRowContext(ctx, `
	   INSERT INTO pastes (user_id, title, content, public_id)
	   VALUES ($1, $2, $3, $4)
	   RETURNING id, user_id, title, content, is_favorite, public_id, created_at, updated_at
	`, userID, title, content, publicID).Scan(
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
		return Paste{}, err
	}

	return paste, nil
}
