// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds 1 billion seconds to the time that's passed
func AddGigasecond(t time.Time) time.Time {
	seconds := math.Pow(10, 9)
	return t.Add(time.Second * time.Duration(seconds))
}
