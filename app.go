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

// Game represents some data that is part of a vidio game
type Game struct {
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Rating    string `json:"rating"`
}

func (g *Game) String() string {
	return fmt.Sprintf("%s %s %s", g.Title, g.Developer, g.Rating)
}

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
	t, err := template.ParseFiles("htmlpages/homepage.html")
	if err != nil {
		fmt.Println("Unable to parse homepage.html", err.Error())
	}
	t.Execute(w, nil)
}

// contribute serves up the contribute page so the client can
// add to the game DB
func contribute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlpages/contribute.html")
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
		fmt.Fprintln(w, "Was unable to submit that game to the DB, try again later")
	} else {
		fmt.Fprintln(w, "That game was successfully submitted")
	}
}

// serveTitleInput will serve up the html page that allows the client
// search the game DB by the games title
func serveTitleInput(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlpages/searchByTitle.html")
	if err != nil {
		fmt.Println("Unable to parse searchByTitle.html")
	}
	t.Execute(w, nil)
}

// serveDeveloperInput will serve up the html page that allows the client
// to search for all games created by the specified developer
func serveDeveloperInput(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlpages/searchByDeveloper.html")
	if err != nil {
		fmt.Println("Unable to parse searchByDeveloper.html")
	}
	t.Execute(w, nil)
}

// serveRatingInput will serve up the html page that allows the client to
// search for all games with the specified rating
func serveRatingInput(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlpages/searchByRating.html")
	if err != nil {
		fmt.Println("Unable to parse searchByRating.html")
	}
	t.Execute(w, nil)
}

// searchByTitle will display the data on the specified game
func searchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	resp, err := http.Get("http://localhost:8080/gameAPI/" + title) // make this better
	if err != nil {
		fmt.Fprintln(w, "Could not connect to the DB, try again later")
		return
	}
	var game Game
	dataLen := resp.ContentLength
	jsonData := make([]byte, dataLen)
	resp.Body.Read(jsonData)
	json.Unmarshal(jsonData, &game)
	fmt.Fprintln(w, fmt.Sprintln("Result:", game.Title, game.Developer, game.Rating)) // out will be better later
}

// searchByDeveloper will display a list of games made by the
// specified developer
func searchByDeveloper(w http.ResponseWriter, r *http.Request) {
	developer := r.FormValue("developer")
	resp, err := http.Get("http://localhost:8080/gameAPI/developer/" + developer)
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
	resp, err := http.Get("http://localhost:8080/gameAPI/rating/" + rating)
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
	t, err := template.ParseFiles("htmlpages/gamelist.html")
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
