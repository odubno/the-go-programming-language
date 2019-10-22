package raindrops

import "fmt"

// Convert converts a number into a sting given some condition
func Convert(n int) string {
	var result string
	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		return fmt.Sprintf("%d", n)
	}
	return result
}
