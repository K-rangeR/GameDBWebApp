package main

// APIError represents the error that is generated when the
// http response from the API is not 404 or 200
type APIError struct{}

func (a APIError) Error() string {
	return "The database is having issues at the moment"
}
