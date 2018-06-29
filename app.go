// gameDBapp is a simple web app that will interact with
// the REST_game api.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/contribute", contribute)
	http.HandleFunc("/search/title", searchByTitle)
	http.HandleFunc("/search/developer", searchByDeveloper)
	http.HandleFunc("/search/rating", searchByRating)
	server.ListenAndServe()
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root")
}

func contribute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contribute")
}

func searchByTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Title")
}

func searchByDeveloper(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Developer")
}

func searchByRating(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rating")
}
