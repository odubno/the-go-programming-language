/*
https://www.youtube.com/watch?v=c78JEGBAopg&list=PLq9Ra239pNZC0MgMN4j6ZiPHv_c0UPnBX&index=10
Defer
	Used to delay execution of a statement until a functioon exists
	Useful to group "open" and "close" functions together
		becareful in loops
	Run in LIFO (last-in, last-out) order
		close open resources in the reverse order that you openned them
	Arguments evaluated at time defer is executed, not at time of called function execution
Panic
	When something unplanned happens in our app there are two things that will happen
		An error will be returned
		Panic will execute
			causes app to shutdown
			used for when an app gets into a state that it can't recover from
			deferred functions will still fire
		If nothing handles panic, program will shutdown
Recover
	built-in
	used to recover from panics
	only useful in deferred functions
	when a panic happens, deferred statements will still execute
		the proper place to use recover is in a deferred function
	current function will not attempt to continue, but higher functions in call stack will
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// deferStatement()
	// deferExample()
	// panicStatement()
	// panicServerExample()
	recoverStatement()
}

func deferStatement() {
	fmt.Println("defer middle statement")
	fmt.Println("start")
	// defer will execute after the function completes its last statement
	// but before the function returns
	defer fmt.Println("middle")
	fmt.Println("end")

	// defer is execute in LIFO (last in last out)
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	defer fmt.Println("defer all statements")
	// we'll use the defer keyword to close out resources
	// it makes sense to close resources in the order they were opened
	// because one resource could be dependent on another

	a := "start"
	defer fmt.Println(a) // start will print out
	// defer takes the arguments at the time defer is called, not when it is executed
	// changing the variable afterwards takes no affect
	a = "end"
}

func deferExample() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	// defer allows to associate the openning and the closing of the same resource
	// useful when there is a lot of code between the openning and closing of a resource
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicStatement() {
	/*
		There are no exceptions in Go
		For example, in Go, if you try to open a file and the file does not exist
			that's a pretty common response. In this case err values are returned.
			Exceptions are not thrown. B/c that is not considered "exceptional" in Go.
		Go uses panic, b/c the application can't function. There's nothing "exceptional"
	*/

	// the runtime will generate a panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")

	// panic happens after deferred statements are executed
	// order:
	// 1. main statements
	// 2. defer (useful for closing out open resources)
	// 3. panic
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}

func panicServerExample() {
	// register a function listener that will listen in on every url in our application
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// call back, when a url comes in that matches "/"
		w.Write([]byte("Hello Go!"))
	})
	// it is reasonable to assume that the http listener could fail
	// the function will not panic it'll return an error
	// We are the ones who will choose how to treat the error
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func recoverStatement() {
	fmt.Println("start")

	// anonymous function does not have a name and is executed exactly one time
	defer func() {
		// this catches any errors and recovers to continue running the rest of the app
		// a warning is raised; the assumption is that you will address the error
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
