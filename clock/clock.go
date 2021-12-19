package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour, min int
}

func New(h, m int) Clock {
	newClock := Clock{}
	minutes := (h * 60) + m

	newClock = newClock.Add(minutes)
	return newClock
}

func (c Clock) Add(m int) Clock {
	if m < 0 {
		return c.Subtract(-m)
	}

	minutes := m % 60
	hours := m / 60

	newMinute := (c.min + minutes)
	if newMinute > 59 {
		newMinute -= 60
		hours++
	}

	newHour := c.hour + hours
	if newHour > 23 {
		newHour %= 24
	}

	return Clock{newHour, newMinute}
}

func (c Clock) Subtract(m int) Clock {
	if m < 0 {
		return c.Add(m)
	}

	minutes := m % 60
	hours := m / 60

	newMinute := c.min - minutes
	if newMinute < 0 {
		newMinute = 60 + newMinute
		hours++
	}

	newHour := c.hour - hours
	if newHour < 0 {
		if newHour%24 == 0 {
			newHour = 0
		} else {
			newHour = 24 + (newHour % 24)
		}
	}

	return Clock{newHour, newMinute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}
