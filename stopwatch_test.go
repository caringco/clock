package clock

import (
	"testing"
)

func TestStopWatch(t *testing.T) {
	sw := NewStopWatch()
	sw.Start()
	sw.Stop()
	sw.TimeElasped()
}
