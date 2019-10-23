package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if any letters repeat in a word
func IsIsogram(word string) bool {
	letters := map[string]int{}
	for _, letter := range strings.ToLower(word) {
		if !unicode.IsLetter(letter) {
			// skip anything that is not a letter
			continue
		}
		if _, ok := letters[string(letter)]; ok {
			// if letter already exists then the word is not an isogram
			return false
		}
		letters[string(letter)] = 1
	}
	return true
}
