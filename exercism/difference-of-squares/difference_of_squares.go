package diffsquares

// SquareOfSum calculates the sum of n numbers and squares the result
func SquareOfSum(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares calculates the square of each n and sums the results
func SumOfSquares(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference calculates the diffence between SumOfSquares and SquareOfSum
func Difference(n int) int {
	return (SumOfSquares(n) - SquareOfSum(n)) * -1
}
