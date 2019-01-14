package clock

import (
	"testing"
	"time"
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
		t.Errorf("expected[%d], actual[%d]", expected, actual)
	}
}

func TestConvertYearsToNano(t *testing.T) {
	years := 7
	expected := int64(years) * Year.Nanoseconds()
	actual := ConvertYearsToNano(years)
	if actual != expected {
		t.Errorf("expected[%d], actual[%d]", expected, actual)
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
	hours := 7
	expected := int64(hours) * time.Hour.Nanoseconds()
	actual := ConvertHoursToNano(hours)
	if actual != expected {
		t.Errorf("expected[%d], actual[%d]", expected, actual)
	}
}

func TestConvertMinToNano(t *testing.T) {
	minutes := 7
	expected := int64(minutes) * time.Minute.Nanoseconds()
	actual := ConvertMinToNano(minutes)
	if actual != expected {
		t.Errorf("expected[%d], actual[%d]", expected, actual)
	}
}

func TestConvertSecToNano(t *testing.T) {
	second := 7
	expected := int64(second) * time.Second.Nanoseconds()
	actual := ConvertSecToNano(second)
	if actual != expected {
		t.Errorf("expected[%d], actual[%d]", expected, actual)
	}
}
