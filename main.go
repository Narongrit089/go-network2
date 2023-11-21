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

	// "/about" is the routh path
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HaHa") // Write data to response
	})

	// static/ is the routh path
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//listen
	http.ListenAndServe(":8080", nil)
}
