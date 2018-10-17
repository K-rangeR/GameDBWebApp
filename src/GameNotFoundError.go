package main

import "fmt"

// GameNotFoundError represents the error that is generated
// when the game the user is looking for was not found by
// the API
type GameNotFoundError struct{}

func (e GameNotFoundError) Error() string {
	return fmt.Sprintf("That game was not found anywhere in the database\n")
}
