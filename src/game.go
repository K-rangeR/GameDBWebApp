// represents a video game with a title, rating, and developer
package main

import "fmt"

// Game represents some data that is part of a vidio game
type Game struct {
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Rating    string `json:"rating"`
}

// String returns a string reprenting the game
func (g *Game) String() string {
	return fmt.Sprintf("%s %s %s", g.Title, g.Developer, g.Rating)
}
