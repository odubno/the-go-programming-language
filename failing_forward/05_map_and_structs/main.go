/*
Maps
	Collections of value types that are accessed by keys
	Created via literals or via make function
	Members accessed via [key] syntax
		myMap["key"] = "value"
	Check for presence with "value, ok" form of result
	Multiple assignments point to the same underlying data

Structs
	collections of disparate data types that describe a single concept
		could point any valid data structure in go
	Keyed by name fields
	Normally created as types, but anonymous structs are allowed
	Structs are value types
	No inheritance, but you can use composition via embedding
	Tags can be added to struct fields to describe field
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Maps
	// Maps()

	// Structs
	//Structs()

	// Embedding and Composition
	Composition()

	// Tags
	Tags()
}

// Maps looks at what maps are
func Maps() {
	/* What are they?
	Creating
	Manipulation
	*/
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	fmt.Println(statePopulations)

	statePopulations2 := make(map[string]int)
	statePopulations2 = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	fmt.Println(statePopulations2)
	fmt.Println(statePopulations2["Texas"])
	statePopulations2["Georgia"] = 10310371
	fmt.Println(statePopulations2["Georgia"])
	// deleting entries
	fmt.Println(statePopulations2)
	delete(statePopulations2, "Georgia")
	fmt.Println(statePopulations2)
	fmt.Println(statePopulations2["Georgia"]) // will return 0
	pop, ok := statePopulations2["Georgia"]
	fmt.Println(pop, ok) // ok will return false
	fmt.Println(len(statePopulations2))
	sp := statePopulations2 // this will create a pointer
	delete(sp, "Texas")     // will delete from statePopulations2 as well
	fmt.Println(sp)
	fmt.Println(statePopulations2)
}

// struct gathers information together that are related to one concept
// we could mix any type of data together
// the struct could get changed without having to change the underlying usage of it
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
	// upper case letters expose the variables outside the package
}

// Structs looks at what structs are
func Structs() {
	/* What are they?
	Creating
	Naming Convention
	Embedding
	Tags
	*/
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Jon Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.Companions[1]) // accessing the slice inside of a struct

	// declare an annonymous struct
	aDoctor2 := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aDoctor2)
	anotherDoctor := aDoctor2
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor2)
	fmt.Println(anotherDoctor)
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal   // embed the Animal struct into the Bird struct
	SpeedKPH float32
	CanFly   bool
}

func Composition() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(c)
}

func Tags() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
