package booking

import (
	"fmt"
	"time"
)

// Default date for Go time parsing: Mon Jan 2 15:04:05 -0700 MST 2006

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// "You have an appointment on Thursday, July 25, 2019, at 13:45."
	layout := "1/2/2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	msg := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		parsed.Weekday(), parsed.Month(), parsed.Day(), parsed.Year(), parsed.Hour(), parsed.Minute())
	return msg
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	now := time.Now()
	anniversary := time.Date(now.Year(), time.September, 15, 00, 00, 00, 00, time.UTC)
	return anniversary
}
