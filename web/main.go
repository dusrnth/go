package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Start the web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":80", nil)
}
