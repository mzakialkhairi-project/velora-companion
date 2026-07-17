package infrastructure

import (
	"errors"
	"time"
)

// Default timezone
const DefaultTimezone = "Asia/Jakarta"

// DefaultLocation returns the default timezone location
func DefaultLocation() *time.Location {
	loc, err := time.LoadLocation(DefaultTimezone)
	if err != nil {
		return time.UTC
	}
	return loc
}

// Now returns current time in default timezone
func Now() time.Time {
	return time.Now().In(DefaultLocation())
}

// ParseTime parses time string with multiple formats
func ParseTime(s string) (time.Time, error) {
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("invalid time format")
}

// FormatTime formats time to RFC3339
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// FormatDate formats time to date only (YYYY-MM-DD)
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats time to datetime (YYYY-MM-DD HH:MM:SS)
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// StartOfDay returns the start of the day
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// AddDays adds days to time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// SubDays subtracts days from time
func SubDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, -days)
}
