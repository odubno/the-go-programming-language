
/*
Variable declaration
	- var foo int
		- useful if you need to declare the variable in a different scope
	- var foo int = 42
		- use when the compiler guesses the wrong type
	- foo := 42
		- will use most frequently
	- can't redeclare variable, but can shadow them
	- all variables must be used

*/


package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Variable Declaration 
	variableDeclaration()

	// Redeclaration and shadowing
	variableRedeclaration()

	// Visibility
		// lower case first letters for package scope
		// upper case first letter to export
		// no privat scope

	// Naming Conventions
		// Lower case variables are scopped only to the package
		// Upper case variables will be expossed to the outside
		// length of the variable complements it's life time
		// variable names are camelCase with acronyms being caps locked
		// PascalCase and camelCase with acronyms always UPPERCASE

	// Type conversions
	typeConversion()
}

// declare variable var4 at the package level
var var4 int = 42

// declare many variables at the package level
var (
	var5 int = 12
	var6 int = 13
	var7 string = "test"
)

func variableDeclaration() {
	// Three different way to declare variables

	// declare a variable var1 of type integer
	var var1 int
	// assign var1 a value
	var1 = 42
	var1 = 27 // var1 value could be changed
	fmt.Println(var1)
	/* 
	sometimes you may want to declare a variable,
	but you're just not ready to use it.
	*/

	// declare a variable var2 of type integer and assign var2 a value
	var var2 int = 42
	fmt.Println(var2)
	/*
	If Go does not have enough information to assign the type you really want.
	If you want Go to assign a specific type, such as int32/int64/float32/float64
	*/

	// go could determine type from provided value
	var3 := 42
	fmt.Println(var3)
	/*
	Useful for when you do not want any specific type.
	*/

	// package level
	fmt.Println(var4)
}

func variableRedeclaration() {
	var var1 int = 42
	var1 = 24
	fmt.Println(var1)
}

func typeConversion() {
	var var1 int = 42
	fmt.Printf("%v, %T\n", var1, var1)

	var var2 float32
	var2 = float32(var1)  // conversion function 
	fmt.Printf("%v, %T\n", var2, var2)

	var var3 string
	var3 = string(var1)
	fmt.Printf("%v, %T\n", var3, var3) // 42 will become *

	// use strconv package for strings
	var3 = strconv.Itoa(var1)
	fmt.Printf("%v, %T\n", var3, var3) // 42 is correctly converted to string
}