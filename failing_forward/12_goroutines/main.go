/*
Creating goroutines
	Use "go" keyword in front of function call
	When using anonymous functions, pass data as a local variable
		avoid using global varaibles in functions
Synchronization
	Use sync.WaitGroup to wait for groups of goroutines to complete
	Use sync.Mutex and sync.RMutex to protext data access.
		one goroutine is manipulating data one at a time
Parallelism
	by default, Go will use CPU threads equal to available cores
		change with runtime.GOMAXPROCS
	More threads can increase performance, but too many can slow it down

Best Practices
	Don't create goroutines in libraries
		let consumer control concurrency
	When creating a goroutine, know how it will end
		avoids subtle memory leaks
	Check for race conditions at compile time
		go run -race main.go
		helps detect the resetting of variables


https://www.youtube.com/watch?v=icbFEmh7Ym0&list=PLq9Ra239pNZC0MgMN4j6ZiPHv_c0UPnBX&index=14
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wait group
// designed to sync multiple go routines together
// this will take just enough time to complete the application
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// 01. Creating goroutines

	// "go" will tell Go to spin up a green thread and run the function in the green thread
	// green threads create an abstraction of a thread referred to as a go routine
	// go sayHello()
	// time.Sleep(100 * time.Millisecond)

	// sayHello2()

	// 02. Synchronization
	// multipleGoRoutinesWaitGroup()
	multipleGoRoutinesMutex()

	// 03. Parallelism
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHello2() {
	var msg = "Hello Again"

	// syncronize to this go routine here
	wg.Add(1)

	go func(msg string) {
		fmt.Println(msg)
		// tell wait group that execution is complete
		// decrement number of tasks the wait group is waiting on
		wg.Done()
	}(msg)

	// once func is done we exit app by waiting
	wg.Wait()
}

func multipleGoRoutinesWaitGroup() {
	// spawning 20 go routines
	// the output will not be printed in order
	for i := 0; i < 10; i++ {

		wg.Add(2) // add 2 go routines

		go sayHello3() // 10 go routines in total
		go increment() // 10 go routines in total
	}
	wg.Wait() // make sure the function does not exit too early
}

func sayHello3() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done() // call done on the wait group
}

func increment() {
	counter++
	wg.Done() // call done on the wait group
}

func multipleGoRoutinesMutex() {
	// using go routines this way; we might as well just remove it all together
	// a lock that the app will honor
	runtime.GOMAXPROCS(100)                          // set the number of threads
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1)) // number of available threads
	// spawning 20 go routines
	for i := 0; i < 10; i++ {

		wg.Add(2) // add 2 go routines

		m.RLock()
		go sayHello4() // 10 go routines in total

		m.Lock()
		go increment4() // 10 go routines in total
	}
	wg.Wait() // make sure the function does not exit too early
}

func sayHello4() {
	fmt.Printf("Hello4 #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment4() {
	counter++
	m.Unlock()
	wg.Done()
}
