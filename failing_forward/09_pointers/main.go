/*
Creating Pointers
	Pointer types use an asterisk(*) as a prefix to type pointed to
		*int - a pointer to an integer
	Use the addressof operator (&) to get address of variable

Dereferencing Pointers
	Dereference a pointer by preceding with an (*)
	Complex types (e.g. structs) are automatically dereferenced

Create pointers to objects
	Can use addressof operator (&) if value type already exists
		ms := *myStruct{foo: 42}
		p := &ms
	Use addressof operator before initializer
		&myStruct{foo: 42}
	Use the new keyword
		Can't initialize fields at the same time

Types with internal pointers
	All assignmentts operations in Go are copy operations
	Slices and maps contain internal pointers, so copies point the same underlying data.

https://www.youtube.com/watch?v=irgo56knnzI&list=PLq9Ra239pNZC0MgMN4j6ZiPHv_c0UPnBX&index=11
*/

package main

import (
	"fmt"
)

func main() {
	// 01. Creating Pointers
	creatingPointers()
	pointerArithmetic()

	// 02. Dereferencing Pointers
	// 03. The new function
	// 04. Working with nil
	// 05. Types with internal pointers
	dereferencingPoiters()
	arrayPointer()
	slicePointer()
	mapPointer()

}

func creatingPointers() {
	a := 42
	b := a
	fmt.Println(a, b)
	a = 25
	// here b does not point to a
	// changing a does not change b
	fmt.Println(a, b)

	var c int = 42
	var d *int = &c
	// & pulls out the memory address of the variable
	// * declares a pointer to an integer
	fmt.Println(c, *d)
	c = 25
	// c and d are tied together
	// both are pointing to the same underlying data
	fmt.Println(c, *d)
}

func pointerArithmetic() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	// can't modify pointers
	// c := &a[1] - 4
	fmt.Printf("%v %p %p\n", a, b, c)
}

type myStruct struct {
	foo int
}

func dereferencingPoiters() {
	var ms *myStruct
	fmt.Println(ms) // zero value for a pointer is <nil>
	ms = &myStruct{foo: 422}
	fmt.Println(ms)
	// need the parantheses to indicate variable is being dereferenced
	(*ms).foo = 42 // same as ms.foo = 42 fmt.Println(ms.foo)
	fmt.Println((*ms).foo)
}

func arrayPointer() {
	// b is a copy of a
	// values of the array are considered intrinsic to the variable
	// a is holding the values of the array as well as the size of the array
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}

func slicePointer() {
	// b points to a
	// slice is a projection of an underlying array
	// slice does not contain the data it self
	// the slice contains a pointer to the first element that
	// the slice is point to in the underlying array
	// internal representation of a slice has a pointer to an array
	// while "a" is copied into "b", part of the data that gets copied is the pointer
	// and not the underlying data it self.
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}

func mapPointer() {
	// similar behavior to slices
	// maps have a pointer to the underlying data
	// maps do not contain the data in them selves
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
	// both values will update

	// primitives, arrays and structs copy the entire structure,
	// unless you're using pointers
}
