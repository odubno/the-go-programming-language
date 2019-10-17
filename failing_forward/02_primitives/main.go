// Primitive Data Types in GO

package main

import (
	"fmt"
)

func main() {
	// Boolean Type
	BooleanType()

	// Numeric Types
	NumericTypes()

	// Text Types
	TextTypes()
}

func BooleanType() {
	/*
		Values are true or false
		Not an alias for other types (e.g. int)
		Zero values is false
	*/

	// useful for state flags
	var var1 bool = true
	fmt.Printf("%v, %T\n", var1, var1)

	// useful for logical tests
	var2 := 1 == 1
	var3 := 2 == 1
	fmt.Printf("%v, %T", var2, var3)
}

func NumericTypes() {
	/*
		Integers
			Signed Integers
				int type has varying size, but min 32 bits
				8 bit (int8) through 64 bit (int64)
				positive and negative numbers
			Unsigned Integers
				8 bit (byte and uint8) through 32 bit (uint32)
				only positive numbers
			Arithmetic operations
				addition, subraction, multiplication, division, remainder
			Bitwise Operations
				and, or, xor and not
			Zero value is 0
			Can't mix types in the same family! (uint32 + unit64)
		Floats
			Follow IEEE-754 standard
			Zero value is 0
			32 and 64 bit versions
			Literal styles
				Decimal(3.14)
				Exponential(13e18 or 2E10)
				Mixed(13.6e10)
			Arithmetic operations
				addition, subtration, multiplication, division
		Comlex Numbers
			Zero value is (0+0i)
			64 and 128 bit versions
			Buit-in functions
				complex
					make complex number from two floats
				real
					get real part as float
				imag
					get imaginary part as float
			Arithmetic operations
				addition, subtration, multiplication, division
	*/
	var1 := 42
	fmt.Printf("%v, %T\n", var1, var1)

	var var2 uint16 = 42
	fmt.Printf("%v, %T\n", var2, var2)

	// Cannot add different int types together

	// bit shifting (only for integers)
	a := 8              // 2^3
	fmt.Println(a << 2) // 2^3 * 2^2 = 2^5 i.e. 32
	fmt.Println(a >> 2) // 2^3 / 2^2 = 2^1 i.e. 2

	// complex numbers
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", real(n), real(n)) // abstracts out the real number
	fmt.Printf("%v, %T\n", imag(n), imag(n)) // abstracts out the imaginary number

	var i complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", i, i)
}

func TextTypes() {
	/*
		Strings
			UTF-8
			Immutable
			Can be concatenated with (+) operator
		Rune
			UTF-32
			Alias for int32
			Special methods require to process
	*/
	// a string in go stands for any utf character
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("strings in go are aliases for bytes: %v, %T\n", s[2], s[2])
	fmt.Printf("convert byte to string representation: %v, %T\n", string(s[2]), string(s[2]))

	// string concat
	fmt.Printf("%v, %T\n", s+"test", s+"test")

	// convert strings to collection of bytes i.e. slices of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// rune represents any utf32 character
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
