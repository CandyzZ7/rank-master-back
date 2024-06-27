package time

import (
	"testing"
	"time"
)

func TestIsToday(t *testing.T) {
	// Example usage
	t1 := time.Now()
	t2 := time.Date(2023, 6, 25, 0, 0, 0, 0, time.Local) // Change to your timezone
	t.Logf("t1 is today:%v", IsToday(t1))
	t.Logf("t2 is today:%v", IsToday(t2))
}
