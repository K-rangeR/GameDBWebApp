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
	resp, err := http.Get(apiTitleEndPoint + title) // test this
	if err != nil {
		fmt.Fprintln(w, "Could not connect to the DB, try again later")
		return
	}
	var game Game
	dataLen := resp.ContentLength
	jsonData := make([]byte, dataLen)
	resp.Body.Read(jsonData)
	json.Unmarshal(jsonData, &game)
	fmt.Fprintln(w, fmt.Sprintln("Result:", game.Title, game.Developer, game.Rating))
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
	var games []Game
	err = unmarshalJSON(&games, resp)
	if err != nil {
		fmt.Println(err.Error())
	}
	parseList(w, games)
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
	var games []Game
	err = unmarshalJSON(&games, resp)
	if err != nil {
		fmt.Println(err.Error())
	}
	parseList(w, games)
}

// unmarshalJSON will convert the game API's json respnse into
// a slice of games
func unmarshalJSON(game *[]Game, r *http.Response) (err error) {
	dataLen := r.ContentLength
	jsonData := make([]byte, dataLen)
	_, err = r.Body.Read(jsonData)
	// if err != nil {
	// 	return
	// }
	err = json.Unmarshal(jsonData, game)
	return
}

// parseList will return an html page containing a list of games to the client
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
