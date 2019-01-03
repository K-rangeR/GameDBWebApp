// gameDBapp is a simple web app that will interact with
// the REST_game api.
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// Relative file paths to the html
const (
	menuPagePath       = "../htmlpages/menu.html"
	homePagePath       = "../htmlpages/homepage.html"
	searchByPath       = "../htmlpages/searchBy.html"
	titleInputPath     = "../htmlpages/searchByTitle.html"
	developerInputPath = "../htmlpages/searchByDeveloper.html"
	ratingInputPath    = "../htmlpages/searchByRating.html"
	bootstrapPath      = "../htmlpages/bootstrapLinks.html"
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

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Could not start the server:", err)
	}
}

// root will serve up the home page
func root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(homePagePath, menuPagePath)
	if err != nil {
		fmt.Println("root:", path.Base(homePagePath))
		return
	}
	t.ExecuteTemplate(w, "home", nil)
}

// serveTitleInput will serve up the html page that allows the client
// search the game DB by using the games title
func serveTitleInput(w http.ResponseWriter, r *http.Request) {
	serveInputPage(w, titleInputPath)
}

// serveDeveloperInput will serve up the html page that allows the client
// to search for all games created by the specified developer
func serveDeveloperInput(w http.ResponseWriter, r *http.Request) {
	serveInputPage(w, developerInputPath)
}

// serveRatingInput will serve up the html page that allows the client to
// search for all games with the specified rating
func serveRatingInput(w http.ResponseWriter, r *http.Request) {
	serveInputPage(w, ratingInputPath)
}

// serveInputPage will serve up the given html page to the client
func serveInputPage(w http.ResponseWriter, pathToPage string) {
	t, err := template.ParseFiles(searchByPath, pathToPage, menuPagePath)
	if err != nil {
		fmt.Println("serverInputPage:", path.Base(pathToPage), err)
		return
	}
	err = t.ExecuteTemplate(w, "searchlayout", nil)
	if err != nil {
		fmt.Println("serverInputPage:", err)
	}
}
