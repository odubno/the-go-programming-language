package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	var sep string
	// os.Args is a slice of strings
	// the first element of os.Args is the path to the file being executed
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}