package paste

import "errors"

var (
	ErrPasteNotFound        = errors.New("paste not found")
	ErrPasteTitleRequired   = errors.New("paste title is required")
	ErrPasteContentRequired = errors.New("paste content is required")
)
