/*
This is an app that simulates a conversation with Bob,
who does not care about much.
*/
package bob

import (
	"strings"
	"unicode"
)

// Hey addresses Bob and always expects a response
func Hey(remark string) string {
	// remove extra spaces
	remark = strings.TrimSpace(remark)

	if isQuestion(remark) && !isUpperCase(remark) {
		// question with no yelling
		return "Sure."
	} else if isUpperCase(remark) && !isQuestion(remark) {
		// yelling with no question
		return "Whoa, chill out!"
	} else if isQuestion(remark) && isUpperCase(remark) {
		// yelling a question
		return "Calm down, I know what I'm doing!"
	} else if isSilence(remark) {
		// address Bob without saying anything
		return "Fine. Be that way!"
	} else {
		// anything else answer whatever.
		return "Whatever."
	}
}

// isUpperCase checks if all chars in string are UPPERCASE and has at least one letter
func isUpperCase(s string) bool {
	for _, v := range s {
		if !unicode.IsLetter(v) {
			// skip any characters that are not letters
			continue
		}
		if !unicode.IsUpper(v) {
			// if even one character is lowercase, return false
			return false
		}
	}
	if !hasLetter(s) {
		// return false is string has no letters
		return false
	}
	return true
}

// isQuestion checks if a string ends in a question mark
func isQuestion(s string) bool {
	if len(s) == 0 {
		return false
	}
	lastCharacter := string(s[len(s)-1:])
	return "?" == lastCharacter
}

func isSilence(s string) bool {
	return "" == strings.TrimSpace(s)
}

func hasLetter(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}
