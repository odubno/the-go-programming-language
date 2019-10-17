/*
Basic Syntax
	Depending on whether the function begins with a lower case or upper case
	determines the visibility of that function
		Upper case first letter -> expose that function outside the package
		Lower case first letter -> keep the function local to the package

Parameters
	provides some variables for the function that are passed from the outside
		comma delimited list of variables and types
			func foo(bar string, baz int)
	Parameters of the same type list type once
		func foo(bar, baz int)
	When pointers are passed in, the function can change the value in the caller
		this is always true for data of slices and maps
			b/c slices and maps work with internal pointers
	Use variadic paremeters to send list of same types
		must be last parameter
		received as a slice
		func foo (bar string, baz ...int)

Return Values
	single return values just like type
		func foo() int
	Multiple return value list types surrounded by parentheses
		func foo() (int, error)
			The (result type, error) paradigm is a very common idiom
	Can use named return values
		initializes returned variables
		return using return keyword on its own
	Can return addresses of local variables
		Automatically promoted from local memory (stack) to shared memory (heap)

Anonymus functions
	Functions don't have names if they are:
		immediately invoked
			func() {
				...
			}()
		assigned to a variable or passed as an argument to a function
			a := func() {
				...
			}
			a()

Functions as types
	can assign functions to variables or use as arguments and return values in functions
	type signature is like function signature, with no parameter names
			var f func(string, string, in) (int, error) // just need to know types comming in and comming out; no arg names

Methods
	a special type of function
	function that executes in context of type
	Format
		// g is the value receiver
		// it means we'll get a copy of the greeter object an it will be passed into the greet method
		func (g greeter) gree() {
			...
		}
	Receiver can be value or pointer
		value receiver gets copy of type
		pointer receiver gets pointer type


https://www.youtube.com/watch?v=jC7iHPedWrM&list=PLq9Ra239pNZC0MgMN4j6ZiPHv_c0UPnBX&index=12
*/

package main

import (
	"fmt"
)

func main() {
	// 01. Parameters
	// parameters()
	// updateVariables()
	// sum(1, 2, 3, 4, 5) // variatic paremeters

	// 02. Return Values
	// returnValues()

	// 03. Anonymous Function & Function as Types
	// anonymousFunc()

	// 04. Methods
	methods()

}

func parameters() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
}

func updateVariables() {
	greeting := "Hello"
	name := "Stacey"
	// passing in the variable reference instead of the content of the variable
	// this will allow us to point to these variables and update them in the function call
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name *string) {
	// by passing in the memory address using (&)
	fmt.Println(*greeting, *name)
	// this will update the memory address with the new value
	*name = "Ted"
	fmt.Println(*name)
}

func sum(values ...int) {
	// variatic parameters
	// the (...) will group all argumentvalues into a slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}

func returnValues() {
	s := sumReturn(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)

	s2 := add(1, 2, 3, 4, 5)
	fmt.Println("The add is", s2)

	d, err := divide(5.0, 1.0)
	// standard test in Go
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func sumReturn(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func add(values ...int) (result int) {
	// named return value
	for _, v := range values {
		result += v
	}
	return // no need to state the return variable; it's implied
}

func divide(a, b float64) (float64, error) {
	// multiple return values
	// one of the values we could return is the error
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero.")
	}
	return a / b, nil
}

// functions them selves can be treated as types, they can be passed around as variables

func anonymousFunc() {
	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	g := func() {
		fmt.Println("Hello Go!")
	}
	f()
	g()
}

func methods() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	// method invocation
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

// method declaration
// a method is a function,
// the difference is the method is executing in a known context
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
