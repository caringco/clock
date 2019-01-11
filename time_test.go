package clock

import (
	"testing"
)

func TestTimeDiff(t *testing.T) {
	wc := NewWallClock()

	wc2 := NewWallClock()
	wc2.MoveClockAt(wc.Time)
	days := 7
	wc2.MoveClockBy(0, days, 0, 0, 0)

	expected := int64(days) * Day.Nanoseconds()
	actual := TimeDiff(wc2.Time, wc.Time)
	if expected != actual {
		t.Error()
	}
}
