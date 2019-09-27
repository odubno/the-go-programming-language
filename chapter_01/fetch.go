package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// os.Args is a slice.
	// it contains the file name as the first index and command line arguments that follow it.
	for _, url := range os.Args[1:] {
		resp, _ := http.Get(url)
		b, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s", b)
	}
}