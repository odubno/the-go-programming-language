package clock

import (
	"fmt"
)

// Clock is a struct
type Clock struct {
	hour, minute int
}

// Normalize normalizes hour and minute
func normalize(hour, minute int) (int, int) {

	// total minutes are remaining
	minutesNormalized := minute % 60

	// total hours are remaining; don't care about days
	hoursNormalized := (hour + minute/60) % 24

	// if remaining hours are negative, add the negative hours to the total hours in day
	if hoursNormalized < 0 {
		hoursNormalized = 24 + hoursNormalized
	}

	// if remaining minutes are negative, add the negaive minutes to the total minutes in one hour
	if minutesNormalized < 0 {
		minutesNormalized = 60 + minutesNormalized
		// when subtracting from minutes, the hour goes down by one
		// before lowering the hour make, sure the value is 24 and not 0
		if hoursNormalized == 0 {
			hoursNormalized = 24
		}
		// lower hours by one
		hoursNormalized--
	}

	return hoursNormalized, minutesNormalized
}

// New is a constructor method for Clock struct
func New(hour, minute int) Clock {
	h, m := normalize(hour, minute)
	return Clock{h, m}
}

// String returns string version of the Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract subtracts minutes from Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
