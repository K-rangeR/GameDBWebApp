// represents a video game with a title, rating, and developer
package main

import "fmt"

// Game represents the parts of a video game that the
// database will track
type Game struct {
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Rating    string `json:"rating"`
}

// String returns a string containing all the game info
func (g *Game) String() string {
	return fmt.Sprintf("%s %s %s", g.Title, g.Developer, g.Rating)
}
