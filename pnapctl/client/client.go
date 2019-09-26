package client

import (
	"io"
	"net/http"
)

// MainClient is the main WebClient used to perform requests.
// Overwrite this variable to change the client used.
var MainClient *http.Client

// BaseURL represents the entry point of our application
var BaseURL string

// PerformGet performs a Get request
func PerformGet(resource string) (*http.Response, error) {
	return MainClient.Get(buildURI(resource))
}

// PerformPost performs a Post request
func PerformPost(resource string, body io.Reader) (*http.Response, error) {
	return MainClient.Post(buildURI(resource), "application/json", body)
}

func buildURI(resource string) string {
	return BaseURL + resource
}
