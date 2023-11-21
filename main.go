package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Hanble request
	r := mux.NewRouter()

	r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		//create cal
		b := []byte("<h1>Book: " + title + "</h1><br><h2>Pade: " + page + "</h2>")

		// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
		w.Write(b)
	})

	// "/" is the routh path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!") // Write data to response
	})

	// "/about" is the routh path
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HaHa") // Write data to response
	})

	// static/ is the routh path
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	//listen
	http.ListenAndServe(":8080", r)
}
