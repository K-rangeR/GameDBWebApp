// gameDBapp is a simple web app that will interact with
// the REST_game api.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Game struct {
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Rating    string `json:"rating"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/contribute", contribute)
	http.HandleFunc("/submit", submitGameToAPI)
	http.HandleFunc("/search/title", searchByTitle)
	http.HandleFunc("/search/developer", searchByDeveloper)
	http.HandleFunc("/search/rating", searchByRating)
	server.ListenAndServe()
}

// root will serve up the home page
func root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		fmt.Println("Unable to parse homepage.html", err.Error())
	}
	t.Execute(w, nil)
}

// contribute serves up the contribute page so the client can
// add to the game DB
func contribute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("contribute.html")
	if err != nil {
		fmt.Println("Unable to parse contribute.html")
	}
	t.Execute(w, nil)
}

// submitGameToAPI will extract the client provided game data
// from the request, convert it into json and send to the game API
func submitGameToAPI(w http.ResponseWriter, r *http.Request) {
	game := Game{r.PostFormValue("title"), r.PostFormValue("developer"), r.PostFormValue("rating")}
	jsonData, err := json.Marshal(game)
	if err != nil {
		fmt.Println("Error creating json")
	}
	_, err = http.Post("http://localhost:8080/gameAPI/add", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Unable to connect to the game API")
	}
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
