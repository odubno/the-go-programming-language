/*
If statements
	initializer syntax
	comparison operators (<, >, <=, >=, ==, !=)
	logical operators (&&, ||, !)
	short circuiting
		if first test evaluates to true, the rest of the staterments will not execute
	if-else statements
	if-else if statements
		additional logical tests
	equality and floats

Switch Statements
	switching on a tag
		a tag is another name for a variable
		test cases are compared against the tag
	cases with multiple tests
	initializers
	switches with no tags
		put in the comparison operator ourselves
			check comparison and logical operators
	fallthrough
		break (implicit)
	type switches
	breaking out early
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// If statements
	// IfStatements()

	// Switch Statements
	SwitchStatements()
}

func IfStatements() {
	// Operators
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	if pop, ok := statePopulations["Florida"]; ok {
		// the pop variable is only defined within the scope of the if statement
		fmt.Println(pop)
	}

	number := 50
	guess := 30

	// test case (base case) "||"" is the "or" operator
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}

	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		// first test
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	fmt.Println(!true)  // false
	fmt.Println(true)   // true
	fmt.Println(!false) // true
	fmt.Println(false)  // false

	// cleaner if statements using else
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		// first test
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	// keep cleaning the if statement
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		// first test
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	// becareful with decimal values
	myNum := 0.12
	myNumModified := math.Pow(math.Sqrt(myNum), 2)
	fmt.Printf("%v == %v\n", myNum, myNumModified)
	if myNum == myNumModified {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

func SwitchStatements() {
	// special purpose if statement
	// first case that passes will be the first one to execute
	// "switch 2", the 2 is a tag when working with switch statements
	// everything is compared against the tag i.e. 2
	switch 2 {
	case 1:
		// 1 is compared against the tag in switch i.e. 2 and if true the conditions will execute
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// compare multiple cases as a range
	switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("this is another number")
	}

	// the tag could be an initializer
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("this is another number")
	}

	// tagless switch statement
	// keep in mind that the break statement is implied
	// first case that meets the condition is executed and the block is exited
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough // keep going; don't break
	case i <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than 20")
	}

	// type switch
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float64")
	case string:
		fmt.Println("j is a string")
	case [3]int:
		fmt.Println("j is an int array of size 3")
	default:
		fmt.Println("j is another type")
	}
}
