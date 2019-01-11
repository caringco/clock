package clock

import (
	"time"
)

const (
	// Day num of nanosec in a day
	Day = 24 * time.Hour
	// Year num of nanosec in a day
	Year = 365 * Day
)

// ConvertDaysToNano convert years to nano
func ConvertDaysToNano(days int) int64 {
	return int64(days) * Day.Nanoseconds()
}

// ConvertHoursToNano convert years to nano
func ConvertHoursToNano(hours int) int64 {
	return int64(hours) * time.Hour.Nanoseconds()
}

// ConvertMinToNano convert years to nano
func ConvertMinToNano(minutes int) int64 {
	return int64(minutes) * time.Minute.Nanoseconds()
}

// ConvertSecToNano convert years to nano
func ConvertSecToNano(seconds int) int64 {
	return int64(seconds) * time.Second.Nanoseconds()
}

// TimeDiff unix nano time difference of two given time
func TimeDiff(x, y time.Time) int64 {
	return x.UnixNano() - y.UnixNano()
}

// ConvertYearsToNano convert years to nano
func ConvertYearsToNano(years int) int64 {
	return int64(years) * Year.Nanoseconds()
}
