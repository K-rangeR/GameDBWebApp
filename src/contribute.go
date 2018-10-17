package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const (
	contributePagePath = "../htmlpages/contribute.html"
	apiAddEndPoint     = "http://localhost:8080/gameAPI/add"
)

// contribute serves up the contribute page so the client can
// add to the game DB
func contribute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(contributePagePath)
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

	resp, err := http.Post(apiAddEndPoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode == http.StatusInternalServerError {
		fmt.Fprintln(w, "Was unable to submit that game to the DB, try again later")
	} else {
		fmt.Fprintln(w, "That game was successfully submitted")
	}
}
