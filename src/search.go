package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const (
	apiTitleEndPoint     = "http://localhost:8080/gameAPI/"
	apiDeveloperEndPoint = "http://localhost:8080/gameAPI/developer/"
	apiRatingEndPoint    = "http://localhost:8080/gameAPI/rating/"
	gameListPagePath     = "../htmlpages/gamelist.html"
)

// searchByTitle will display the data on the specified game
func searchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	resp, err := http.Get(apiTitleEndPoint + title)
	if err != nil {
		fmt.Fprintln(w, "Could not connect to the DB, try again later")
		return
	}
	if searchSuccessful(resp, w) {
		var game Game
		dataLen := resp.ContentLength
		jsonData := make([]byte, dataLen)
		resp.Body.Read(jsonData)
		json.Unmarshal(jsonData, &game)
		fmt.Fprintln(w, fmt.Sprintln("Result:", game.Title, game.Developer, game.Rating))
	}
}

// searchByDeveloper will display a list of games made by the
// specified developer
func searchByDeveloper(w http.ResponseWriter, r *http.Request) {
	developer := r.FormValue("developer")
	resp, err := http.Get(apiDeveloperEndPoint + developer)
	if err != nil {
		fmt.Fprintln(w, "Could not connect to the DB, try again later")
		return
	}
	if searchSuccessful(resp, w) {
		var games []Game
		err = unmarshalJSON(&games, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		parseList(w, games)
	}
}

// searchByRating will display a list of all games with the
// specified rating
func searchByRating(w http.ResponseWriter, r *http.Request) {
	rating := r.FormValue("rating")
	resp, err := http.Get(apiRatingEndPoint + rating)
	if err != nil {
		fmt.Fprintln(w, "Could not connect to the DB, try again later")
		return
	}
	if searchSuccessful(resp, w) {
		var games []Game
		err = unmarshalJSON(&games, resp)
		if err != nil {
			fmt.Println(err.Error())
		}
		parseList(w, games)
	}
}

// unmarshalJSON will convert the game API's json respnse into
// a slice of games
func unmarshalJSON(game *[]Game, r *http.Response) (err error) {
	dataLen := r.ContentLength
	jsonData := make([]byte, dataLen)
	_, err = r.Body.Read(jsonData)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonData, game)
	return
}

// parseList will send an html page containing a list of games to the client
func parseList(w http.ResponseWriter, games []Game) {
	output := getGameListOutput(games)
	t, err := template.ParseFiles(gameListPagePath)
	if err != nil {
		fmt.Println("Error parsing game list")
	}
	t.Execute(w, output)
}

// getGameListOutput returns a slice of strings where each
// string represents a game from the games parameter
func getGameListOutput(games []Game) []string {
	gamesAsString := make([]string, 0)
	for _, game := range games {
		gamesAsString = append(gamesAsString, game.String())
	}
	return gamesAsString
}

// searchSuccessful will check the status code of the game API response to
// determine if the game was successfully located. Returns true if status
// code is 200, and displays an error message and returns false otherwise
func searchSuccessful(resp *http.Response, w http.ResponseWriter) bool {
	fmt.Println("Code:", resp.StatusCode)
	if resp.StatusCode == 404 {
		fmt.Fprintln(w, "Could not find that game anywhere in the database")
		return false
	} else if resp.StatusCode != 200 {
		fmt.Fprintln(w, "The database is having issues at the moment")
		return false
	} else {
		return true
	}
}
