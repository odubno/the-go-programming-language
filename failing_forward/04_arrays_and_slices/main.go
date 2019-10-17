package main

import (
	"fmt"
)

func main() {
	// Arrays
	Arrays()

	// Slices
	Slices()
}

func Arrays() {
	// Arrays form the basis for slices

	// 01. Creation
	grades := [3]int{97, 85, 93} // An array could only hold one type
	fmt.Printf("Grades: %v\n", grades)

	// If you're initializing an array literal then no need to indicate it's size
	// the size is infered from the data that is passed in
	// the ... tell go to create an array that is just large enough to hold the data that is passed in
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades2)

	var students [3]string // initialize an array
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa" // add value to array
	fmt.Printf("Students: %v\n", students)

	// identity matrix
	// array that stores multiple array
	// matrix 3x3
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println(identityMatrix)

	// an easier way to see the above matrix
	var identityMatrix2 [3][3]int
	// initialize rows individually
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)

	// Arrays in Go are considered to be values
	// in other languages when you create, it's pointing to the values in the array
	a := [...]int{1, 2, 3}
	// this will not point to the same array a, this will create a copy
	b := a
	b[1] = 5
	fmt.Println(a, b)
	c := &a // c will point to the same data as a
	c[1] = 10
	fmt.Println(a, c)

	// 02. Built-in functions

	// 03. Working with arrays
}

func Slices() {
	// 01. Creation
	a := []int{1, 2, 3}
	fmt.Printf("%v\n", a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// slices are natually reference types and point to the same data when assigned to new variables
	// unlike arrays, we do not need to explicitly point to a memory address
	b := a
	b[1] = 5
	fmt.Println(a, b)

	// 02. Built-in functions
	// make(type, size, capacity)
	c := make([]int, 3, 100)
	fmt.Printf("%v\n", c)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))

	// 03. Working with slices
}
