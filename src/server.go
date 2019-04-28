package main

import (
	"fmt"
	"net/http"
	"time"
)

// request handler
func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func second(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World suckers!")
	// http.Redirect(w, r, "/", http.StatusFound)
}

func third(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World I'm at %s", request.URL.Path[1:])
}

func main() {
	// routes
	http.HandleFunc("/", greet)
	http.HandleFunc("/new", second)
	http.HandleFunc("/hmm", third)
	// start the server and listen for requests
	http.ListenAndServe(":8080", nil)
}
