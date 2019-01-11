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

func TestConvertDaysToNano(t *testing.T) {
	days := 7
	expected := int64(days) * Day.Nanoseconds()
	actual := ConvertDaysToNano(days)
	if actual != expected {
		t.Errorf("expected[%d], actual[%d]", expected, actual)
	}
}

func TestConvertHoursToNano(t *testing.T) {
}

func TestConvertMinToNano(t *testing.T) {
}

func TestConvertSecToNano(t *testing.T) {
}

func TestConvertYearsToNano(t *testing.T) {
}
