package time

import (
	"time"
)

const (
	StandardDateFormat = "20060102"
)

// IsToday checks if the given time is today
func IsToday(t time.Time) bool {
	now := time.Now()
	year, month, day := now.Date()
	targetYear, targetMonth, targetDay := t.Date()
	return year == targetYear && month == targetMonth && day == targetDay
}
