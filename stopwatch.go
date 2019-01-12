package clock

import (
	"fmt"
	"time"
)

// StopWatch data structure
type StopWatch struct {
	StartTime time.Time
	EndTime   time.Time
	// delta useful in implementing Pause
	delta int64
}

// NewStopWatch returns new instance of StopWatch
//
func NewStopWatch() *StopWatch {
	return &StopWatch{StartTime: time.Now()}
}

// Reset resets the stopwatch StartTime, EndTime to time.Now()
func (s *StopWatch) Reset() {
	s.StartTime = time.Now()
	s.EndTime = time.Now()
}

// Start starts the stopwatch
func (s *StopWatch) Start() {
	s.Reset()
}

// Pause pauses the time
//
// This feature is currenlty in work in progress state
//
func (s *StopWatch) Pause() {
	s.EndTime = time.Now()
	s.delta += TimeDiff(s.EndTime, s.StartTime)
	s.Reset()
}

// Stop stops the stopwatch
func (s *StopWatch) Stop() {
	s.EndTime = time.Now()
}

// TimeElasped between start and stop
func (s *StopWatch) TimeElasped() int64 {
	d := TimeDiff(s.EndTime, s.StartTime)
	return s.delta + d
}

// ToString string representation of stopwatch
func (s *StopWatch) String() string {
	return fmt.Sprintf("StartTime: %s, EndTime: %s", s.StartTime.Format("2006-01-02 15:04:05.000000000"), s.EndTime.Format("2006-01-02 15:04:05.000000000"))
}
