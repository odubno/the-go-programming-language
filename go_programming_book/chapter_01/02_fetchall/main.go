/*
Fetch URLs concurrently using goroutines and channels

go build 02_fetchall.go
./02_fetchall https://golang.org http://gopl.io https://godoc.org
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// a goroutine is a concurrent function execution
	// a channel allows one goroutine to pass values to another goroutine
	start := time.Now()

	// creates a channel of strings using make
	ch := make(chan string)

	urls := os.Args[1:]
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch and print the lines
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// Copy returns the byte count, along with any error that occurred
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
