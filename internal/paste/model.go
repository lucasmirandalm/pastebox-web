package paste

import "time"

type Paste struct {
	ID         int64
	UserID     int64
	Title      string
	Content    string
	IsFavorite bool
	PublicID   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
