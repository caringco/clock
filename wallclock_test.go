package clock

import (
	"testing"
	"time"
)

func TestMoveClockAt(t *testing.T) {
	wc := NewWallClock()
	givenTime := time.Now()
	wc.MoveClockAt(givenTime)
	if givenTime != wc.GetTime() {
		t.Errorf("expected[%s], actual[%s]", givenTime.String(), wc.GetTime().String())
	}
}

func TestMoveClockBy(t *testing.T) {
	wc := NewWallClock()
	days := 7
	expectedTime := wc.GetTime().UnixNano() + ConvertDaysToNano(days)
	wc.MoveClockBy(0, days, 0, 0, 0)
	actualTime := wc.GetTime().UnixNano()
	if expectedTime != actualTime {
		t.Errorf("expected[%d], actual[%d]", expectedTime, actualTime)
	}
}
