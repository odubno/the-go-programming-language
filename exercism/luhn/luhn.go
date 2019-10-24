package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid confirms whether a number is valid luhn number
// double every second digit, starting from the right
// if the sum is greater than 9, subtract 9 from the sum
func Valid(num string) bool {
	var result int
	num = strings.TrimSpace(num)

	if len(num) < 2 {
		return false
	}

	// toDouble tracks when a number should be doubled
	var toDouble bool

	// start from the right of the num string
	for i := len(num) - 1; i >= 0; i-- {

		// if a character is a space, skip it
		if unicode.IsSpace(rune(num[i])) {
			continue
		}

		// if a character is not a digit the number is invalid
		if !unicode.IsDigit(rune(num[i])) {
			return false
		}

		// convert string character to integer
		n, _ := strconv.Atoi(string(num[i]))

		if toDouble {
			n += n
			if n > 9 {
				n -= 9
			}
			toDouble = false
		} else {
			toDouble = true
		}

		result += n
	}

	// if the result is evenly divisible by 10, then the number is valid
	return result%10 == 0
}
