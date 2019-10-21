package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate converts a phrase to an acronym
func Abbreviate(s string) string {
	// split the string on various characters
	re := regexp.MustCompile(" _|_ | - |-| ")
	words := re.Split(s, -1)
	var acronym string
	for _, word := range words {
		firstLetter := string(word[0])
		acronym += firstLetter
	}
	return strings.ToUpper(acronym)
}
