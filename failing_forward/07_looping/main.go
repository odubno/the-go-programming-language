/*
For Staterments
	simple loops
		for initializer; test; incrementer {}
		for test {}
		for {}
			will run indefinetely
	exiting early
		break
			breaks out of the current iteration
		continue
		labels
			breaks out of the outter loop of a nested loop
	looping over collections
		arrays, slices, maps, strings, channels
		for k, v := range collection {}
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("initializer; bool result; incrementer")
	for i := 0; i < 5; i++ {
		// i is scoped to the for block
		fmt.Println(i)
	}

	fmt.Println("loop over multiple variables")
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("initialize i outside the for loop")
	i := 0 // i is scopped to the main function
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("increment j inside the for loop")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("while loop")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 6 {
			break
		}
	}

	fmt.Println("while loop with continue staterment")
	l := 0
	for {
		l++
		if l == 3 {
			continue
		}
		fmt.Println(l)
		if l == 6 {
			break
		}
	}

	fmt.Println("Break out of the outer loop using a Label")
MyLoopLabel:
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 3; n++ {
			fmt.Println(m * n)
			if m*n >= 3 {
				break MyLoopLabel
			}
		}
	}

	fmt.Println("Working with collections and for loops")
	fmt.Println("using for range loop")
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println("Accessing maps in for loops")
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	for _, v := range statePopulations {
		fmt.Println(v)
	}

}
