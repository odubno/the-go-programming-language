package hamming

import "fmt"

// Distance calculates the Hamming distance between two DNA strands.
// Keep count of the letters in both strings that are different.
func Distance(a, b string) (int, error) {
	var count int
	if len(a) != len(b) {
		return count, fmt.Errorf("length of both strings must match")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
