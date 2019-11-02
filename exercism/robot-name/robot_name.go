package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot is pointer type for the name of the Robot
type Robot string

var robotNames []string

func init() {
	// Seed the random number generator using the current time (nanoseconds since epoch)
	rand.Seed(time.Now().UTC().UnixNano())
}

// Name returns the existing Robot name
func (r *Robot) Name() (string, error) {
	// If a robot name does not exist, define one
	if *r == "" {
		r.Reset()
	}
	if len(robotNames) == 0 {
		// after Robot is instantiated, it's being called more than 676,000 times
		return "", fmt.Errorf("namespace has been exhausted")
	}
	return string(*r), nil
}

// Reset overrides existing robot name
func (r *Robot) Reset() {
	if len(robotNames) == 0 {
		// All possible 2 letter and 3 digit combinations
		robotNames = getRobotNames()
	}
	// reset Robot name
	robotName := generateName()
	*r = Robot(robotName)
}

// generateName returns a random robot name
func generateName() string {
	// index out random robot name
	idx := rand.Intn(len(robotNames))
	robotName := string(robotNames[idx])
	robotNames = pop(robotNames, idx)
	return robotName
}

// combination creates n length combo of (s) string
func combination(s string, prefix string, length int, res *[]string) {
	if length == 0 {
		*res = append(*res, prefix)
		return
	}
	for i := 0; i < len(s); i++ {
		newPrefix := prefix + string(s[i])
		combination(s, newPrefix, length-1, res)
	}
	return
}

// pop removes an element at index i from a
func pop(a []string, i int) []string {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.
	return a
}

// getRobotNames creates all possible combinations of letters and digits
func getRobotNames() []string {
	// 2 value combination of letters
	var randomLetters []string
	combination("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "", 2, &randomLetters)
	// 3 value combination of digits
	var randomDigits []string
	combination("1234567890", "", 3, &randomDigits)
	var names []string
	for l := range randomLetters {
		letter := string(randomLetters[l])
		for d := range randomDigits {
			digit := string(randomDigits[d])
			names = append(names, letter+digit)
		}
	}
	return names
}
