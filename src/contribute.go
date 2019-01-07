package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const (
	contributePagePath    = "../htmlpages/contribute.html"
	successfulPagePath    = "../htmlpages/successfulAdd.html"
	internalErrorPagePath = "../htmlpages/internalError.html"
	apiAddEndPoint        = "http://localhost:8080/gameAPI/add"
)

// contribute serves up the contribute page so the client can
// add a game to the game DB
func contribute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(contributePagePath, menuPagePath, bootstrapPath)
	if err != nil {
		fmt.Println("Unable to parse contribute.html:", err)
		return
	}
	t.ExecuteTemplate(w, "contribute", nil)
}

// submitGameToAPI will extract the client provided game data
// from the request, convert it into json and send it to the game API
func submitGameToAPI(w http.ResponseWriter, r *http.Request) {
	game := Game{r.PostFormValue("title"), r.PostFormValue("developer"), r.PostFormValue("rating")}
	jsonData, err := json.Marshal(game)
	if err != nil {
		sendErrorHTML(w)
		return
	}

	resp, err := http.Post(apiAddEndPoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		sendErrorHTML(w)
	} else {
		sendSuccessHTML(w)
	}
}

// sendSuccessHTML sends the successful add html page
func sendSuccessHTML(w http.ResponseWriter) {
	t, err := template.ParseFiles(successfulPagePath, menuPagePath, bootstrapPath)
	if err != nil {
		fmt.Println("sendSuccessHTML:", err)
		fmt.Fprintln(w, "That game has been added to the database!")
	}
	t.ExecuteTemplate(w, "success", nil)
}

// sendErrorHTML sends the error reporting html page to the client
func sendErrorHTML(w http.ResponseWriter) {
	t, err := template.ParseFiles(internalErrorPagePath, menuPagePath, bootstrapPath)
	if err != nil {
		fmt.Println("sendErrorHTML:", err)
		fmt.Fprintln(w, "Oops, looks like there was an error, please try again later")
	}
	t.ExecuteTemplate(w, "error", nil)
}
