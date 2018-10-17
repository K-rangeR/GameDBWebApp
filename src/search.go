package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
	if err = searchSuccessful(resp.StatusCode); err == nil {
		var game Game
		dataLen := resp.ContentLength
		jsonData := make([]byte, dataLen)
		resp.Body.Read(jsonData)
		json.Unmarshal(jsonData, &game)
		fmt.Fprintln(w, fmt.Sprintln("Result:", game.Title, game.Developer, game.Rating))
	} else {
		reportError(w, err)
	}
}

// searchByDeveloper will display a list of games made by the
// specified developer
func searchByDeveloper(w http.ResponseWriter, r *http.Request) {
	developer := r.FormValue("developer")
	games, err := searchBy(apiDeveloperEndPoint, developer)
	if err != nil {
		reportError(w, err)
		return
	}
	parseList(w, games)
}

// searchByRating will display a list of all games with the
// specified rating
func searchByRating(w http.ResponseWriter, r *http.Request) {
	rating := r.FormValue("rating")
	games, err := searchBy(apiRatingEndPoint, rating)
	if err != nil {
		reportError(w, err)
		return
	}
	parseList(w, games)
}

// searchBy will send an http GET request to the API in search
// for games that match some criteria (by)
func searchBy(endPoint, by string) ([]Game, error) {
	resp, err := http.Get(endPoint + by)
	if err != nil {
		return nil, err
	}

	if err = searchSuccessful(resp.StatusCode); err != nil {
		return nil, err
	}

	var games []Game
	err = unmarshalJSON(&games, resp)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return games, nil
}

// searchSuccessful will check the status code of the game API response to
// determine if the game was successfully located
func searchSuccessful(statusCode int) error {
	// replace with switch
	if statusCode == http.StatusNotFound {
		fmt.Println("games not found")
		return fmt.Errorf("that game was not found anywhere is the database")
	} else if statusCode != http.StatusOK {
		fmt.Println("other error")
		return fmt.Errorf("the database is having an issue at the moment")
	} else {
		fmt.Println("games found")
		return nil
	}
}

// unmarshalJSON will convert the game API's json respnse into
// a slice of games
func unmarshalJSON(game *[]Game, r *http.Response) (err error) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error in unmarshalJSON")
		return
	}
	err = json.Unmarshal(jsonData, game)
	return
}

// reportError will inform the client on the type of error that has occured
func reportError(w http.ResponseWriter, err error) {
	fmt.Fprintln(w, err.Error())
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
