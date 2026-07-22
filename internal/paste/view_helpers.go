package paste

import (
	"fmt"
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
