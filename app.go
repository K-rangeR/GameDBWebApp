// gameDBapp is a simple web app that will interact with
// the REST_game api.
package main

import (
	"fmt"
	"html/template"
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

// root will serve up the home page
func root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		fmt.Println("Unable to pare homepage.html", err.Error())
	}
	t.Execute(w, nil)
}

// contribute serves up the contribute page
func contribute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contribute")
}

// searchByTitle will display the data on the specified game
func searchByTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Title")
}

// searchByDeveloper will display a list of games made by the
// specified developer
func searchByDeveloper(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Developer")
}

// searchByRating will display a list of all games with the
// specified rating
func searchByRating(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rating")
}
