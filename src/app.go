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
	http.HandleFunc("/submit", submitGameToAPI)
	http.HandleFunc("/getgame/title", serveTitleInput)
	http.HandleFunc("/getgame/developer", serveDeveloperInput)
	http.HandleFunc("/getgame/rating", serveRatingInput)
	http.HandleFunc("/search/title", searchByTitle)
	http.HandleFunc("/search/developer", searchByDeveloper)
	http.HandleFunc("/search/rating", searchByRating)
	server.ListenAndServe()
}

// root will serve up the home page
func root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../htmlpages/homepage.html")
	if err != nil {
		fmt.Println("Unable to parse homepage.html", err.Error())
	}
	t.Execute(w, nil)
}

// serveTitleInput will serve up the html page that allows the client
// search the game DB by the games title
func serveTitleInput(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../htmlpages/searchByTitle.html")
	if err != nil {
		fmt.Println("Unable to parse searchByTitle.html")
	}
	t.Execute(w, nil)
}

// serveDeveloperInput will serve up the html page that allows the client
// to search for all games created by the specified developer
func serveDeveloperInput(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../htmlpages/searchByDeveloper.html")
	if err != nil {
		fmt.Println("Unable to parse searchByDeveloper.html")
	}
	t.Execute(w, nil)
}

// serveRatingInput will serve up the html page that allows the client to
// search for all games with the specified rating
func serveRatingInput(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../htmlpages/searchByRating.html")
	if err != nil {
		fmt.Println("Unable to parse searchByRating.html")
	}
	t.Execute(w, nil)
}
