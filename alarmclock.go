// @wip (work in progress)
package clock

import (
	"time"
)

// AlarmClock data structure
type AlarmClock struct {
	time time.Time
	at   time.Time
}

// NewAlarmClock creates new alarmclock instance
func NewAlarmClock() *AlarmClock {
	return &AlarmClock{}
}

// SetAlarmAt set alarm at the given time
func (a *AlarmClock) SetAlarmAt(t time.Time) {
	a.at = t
}
