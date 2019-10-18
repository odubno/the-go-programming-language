# the-go-programming-language

This is my sand box for learning and experimenting with Go. Each chapter directory is from `The Go Programming Language` by Alan A. A. Donovan and Brian W. Kerninghan. Below, you'll find some notes that I wrote down as I went through each chapter. 

The `failing_forward` directory contains work from tutorials made by Mike Van Sickle [Failing Forward - YouTube](https://www.youtube.com/watch?v=OSPNUKoN81o&list=PLq9Ra239pNZC0MgMN4j6ZiPHv_c0UPnBX). It goes over the workflow using Go and its tools. Each subdirectory has a `main.go`; there you'll find my code and notes.

The `exercism` directory contains work from excercises found on [exercism.io](https://exercism.io/profiles/odubno).

## chapter 1

Go is a compiled language. 

`go run helloworld.go` 
- `go` is a terminal command that has some subcommands, and `run` is one of them. 
- `run` compiles the source code and links it with libraries, then runs the resulting executable file.

Go natively handles Unicode and could easily process other world languages. 

`go build helloworld.go`
- `build` compiles the program once
- an executable binary file is generated called `helloworld` that can be run using `./helloworld`.

Go is indexed at 0.

### for loop
- all statements in the for loop are optional. If none are present, the for loop is infinite. 
```golang
for initialization; condition; post {
    // zero or more statements 
}
```
- **initialization** is executed before the loop starts. Whern present, this must be a simple statement.
    - variable declaration; increment; assignment; function call
- **condition** is a boolean expression that is evaluated at the beginning of each iteration of the loop
- **post** is executed after the body of the loop, then the condition is evaluated again. The loop ends when the condition becomes false.

### Understanding Types in Go
- https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html
- Type is Life.
- Type is everything.
- Type defines how the bits gets used.
- Type determines the amount of memory, in bytes, to look at

#### Unsigned Integers
- uint8, uint16, uint32, uint64

#### Signed Integers
- int8, int16, int32, int64

#### Real Numbers
- float32, float64

#### Predeclared Integers
- uint, int, uintptr

## chapter 2

Going into the basic structure 