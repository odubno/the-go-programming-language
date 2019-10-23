package reverse

// Reverse revereses the order of a given string
func Reverse(s string) string {
	// convert string to a list of runes;
	// a rune stores unicode representation of string characters
	data := []rune(s)
	var result string

	// execute an inverse loop over data
	for i := len(data) - 1; i >= 0; i-- {
		// convert each rune to string and append it in reverse order
		result += string(data[i])
	}
	return result
}
