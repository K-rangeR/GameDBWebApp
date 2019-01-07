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
	gameNotFoundPagePath = "../htmlpages/gameNotFound.html"
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
		parseList(w, []Game{game})
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
	switch statusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return fmt.Errorf("that game was not found anywhere in the database")
	default:
		return fmt.Errorf("the databasd is having an issue at the moment")
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
	t, err := template.ParseFiles(gameNotFoundPagePath, menuPagePath, bootstrapPath)
	if err != nil {
		fmt.Println("reportError:", err)
		fmt.Fprintln(w, "Looks like there are no games that match that criteria")
		return
	}
	t.ExecuteTemplate(w, "notfound", nil)
}

// parseList will send an html page containing a list of games to the client
func parseList(w http.ResponseWriter, games []Game) {
	t, err := template.ParseFiles(gameListPagePath, menuPagePath, bootstrapPath)
	if err != nil {
		fmt.Println("Error parsing game list")
	}
	t.ExecuteTemplate(w, "gameslist", games)
}
