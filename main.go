package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "/" is the routh path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!") // Write data to response
	})

	//listen
	http.ListenAndServe(":8080", nil)
}
