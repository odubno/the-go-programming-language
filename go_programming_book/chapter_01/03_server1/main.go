// Server1 is a minimal "echo" server
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
}
