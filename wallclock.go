package clock

import (
	"time"
)

// WallClock data structure
type WallClock struct {
	Time time.Time
}

// NewWallClock create new wallclock object with time now
func NewWallClock() *WallClock {
	return &WallClock{
		Time: time.Now(),
	}
}

// Reset wallclock time to time now
func (w *WallClock) Reset() {
	w.Time = time.Now()
}

// MoveClockAt the given time
func (w *WallClock) MoveClockAt(t time.Time) {
	w.Time = t
}

// AddDate add number of years, months, days
func (w *WallClock) AddDate(years int, months int, days int) {
	w.Time.AddDate(years, months, days)
}

// MoveClockBy number of days, hours, min, sec
func (w *WallClock) MoveClockBy(years, days, hours, min, sec int) {
	tNano := ConvertYearsToNano(years) + ConvertDaysToNano(days) + ConvertHoursToNano(hours) + ConvertMinToNano(min) + ConvertSecToNano(sec)
	var d = time.Duration(tNano)
	w.MoveClockAt(w.Time.Add(d))
}

// GetTime is the current time of wallckock
func (w *WallClock) GetTime() time.Time {
	return w.Time
}

// GetUnixNano representation of wallclock
func (w *WallClock) GetUnixNano() int64 {
	return w.Time.UnixNano()
}
