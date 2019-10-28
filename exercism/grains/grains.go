package grains

import (
	"fmt"
)

// Square returns the amount of grain on a given square
func Square(num int) (uint64, error) {
	if num > 64 || num < 1 {
		return 0, fmt.Errorf("only numbers including and between 1 and 64 are allowed")
	}
	var result uint64
	result = 1
	for i := 1; num >= i+1; i++ {
		result += result
	}
	return result, nil
	// fast solution
	// return uint64(math.Pow(2.0, float64(num-1))), nil
	// even better and faster solution
	//return 1 << (num - 1), nil
}

// Total returns the total grainss accross 64 squares
func Total() uint64 {
	var result uint64
	for i := 1; 64 >= i; i++ {
		square, _ := Square(i)
		result += square
	}
	return result
}
