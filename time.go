package clock

import (
	"time"
)

const (
	// Day num of nanosec in a day
	Day = 24 * time.Hour
	// Year num of nanosec in a year
	Year = 365 * Day
)

// TimeFunc provides the current system time by default
//
// You can override it to use another time value.  This is useful for testing or if you
// want to use a different time zone than your system.
var TimeFunc = time.Now

// ConvertYearsToNano convert years to nanoseconds
func ConvertYearsToNano(years int) int64 {
	return int64(years) * Year.Nanoseconds()
}

// ConvertDaysToNano convert days to nanoseconds
func ConvertDaysToNano(days int) int64 {
	return int64(days) * Day.Nanoseconds()
}

// ConvertHoursToNano convert hours to nanoseconds
func ConvertHoursToNano(hours int) int64 {
	return int64(hours) * time.Hour.Nanoseconds()
}

// ConvertMinToNano convert minutes to nanoseconds
func ConvertMinToNano(minutes int) int64 {
	return int64(minutes) * time.Minute.Nanoseconds()
}

// ConvertSecToNano convert seconds to nanoseconds
func ConvertSecToNano(seconds int) int64 {
	return int64(seconds) * time.Second.Nanoseconds()
}

// TimeDiff unix nano time difference of two given time
func TimeDiff(x, y time.Time) int64 {
	return x.UnixNano() - y.UnixNano()
}
