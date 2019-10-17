/*
Basics

Composing interfaces

Type Conversion
	The empty interface
	Type switches

Implementing with values vs pointers

Best Practices
*/

package main

import (
	"fmt"
)

func main() {
	// 01. Basics
	basics()
}

func basics() {
	// first example
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	// second example
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

}

type Writer interface {
	// structs are data containers; structs describe data
	// interfaces describe behaviors and store method definitions
	// instead of entering a bunch of data types
	// we store method definitions

	Write([]byte) (int, error)
}

type ConsoleWriter struct {
	// interfaces are not implemented explicitly
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	a, err := fmt.Println(string(data))
	return a, err
}

type Incrementer interface {
	Increment() int
}

// type alias for an interger
type IntCounter int

// add a method to the custom IntCounter
func (ic *IntCounter) Increment() int {
	*ic++ // incrementing the actual type
	return int(*ic)
}
