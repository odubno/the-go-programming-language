// Constants are immutable but can be shadowed
// Replaced by the compiler at compile time
// value must be calculated at compile time
// Named like variables
// PascalCase for exported constants
// camelCase for internal constants
// Typed constants work like immutable variables
// Can interoperate only with same type
// Untyped constants work like literals
// Can interoperate with similar types

package main

import (
	"fmt"
)

func main() {
	// Naming convention
	NamingConvention()

	// Typed constants
	TypedConstansts()

	// Untyped constants
	UntypedConstants()

	// Enumerated constants
	EnumeratedConstants()

	// Enumeration expressions
	EnumerationExpressions()

}

func NamingConvention() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
}

// gets overwritten TypedConstansts()
const a int16 = 27

func TypedConstansts() {
	// can't set constants to something that has to be determined at run time
	// var var1 float64 = math.Sin(1.57)
	// const myConst2 float64 = var1

	const a int = 14 // overwrittes package level constant
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	var e int = 27
	fmt.Printf("%v, %T\n", a+e, a+e)
}

func UntypedConstants() {
	// we could let go infer the type
	const f = 42
	fmt.Printf("%v, %T\n", f, f)

	// go infers the "f" value
	// the compiler sees "f" as a literal 42 and is interpreted as an int16
	// go will implicitly convert the value
	var g int16 = 27
	fmt.Printf("%v, %T\n", f+g, f+g)
}

const h = iota // counter for enumerated constanst

// patern translates to the rest of the variables
const (
	i = iota
	j
	k
)

func EnumeratedConstants() {
	// iota is scoped to a constant block and will reset
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
}

func EnumerationExpressions() {
	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials
		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
}
