// Exercise 1.12: Modify the Lissajous server to read parameter values
// from the URL. For example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead
// of the default 5. Use the strconv.Atoi function to convert the string
// parameter into an integer. You can see its documentation by running
// go doc strconv.Atoi

package main

import (
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// please note that HandleFunc run in goroutine
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
