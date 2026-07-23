package paste

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func formatLastEdited(t time.Time) string {
	duration := time.Since(t)

	if duration < time.Minute {
		return "now"
	}

	if duration < time.Hour {
		minutes := int(duration.Minutes())

		if minutes == 1 {
			return "1 minute ago"
		}

		return fmt.Sprintf("%d minutes ago", minutes)
	}

	if duration < 24*time.Hour {
		hours := int(duration.Hours())

		if hours == 1 {
			return "1 hour ago"
		}

		return fmt.Sprintf("%d hours ago", hours)
	}

	if duration < 48*time.Hour {
		return "yesterday"
	}

	days := int(duration.Hours() / 24)

	return fmt.Sprintf("%d days ago", days)
}

func buildPastesURL(onlyFavorites bool, search string) string {
	values := url.Values{}

	if onlyFavorites {
		values.Set("filter", "favorites")
	}

	search = strings.TrimSpace(search)
	if search != "" {
		values.Set("q", search)
	}

	query := values.Encode()
	if query == "" {
		return "/pastes"
	}

	return "/pastes?" + query
}
