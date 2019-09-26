package client

import (
	"io"
	"net/http"
)

// MainClient is the main WebClient used to perform requests.
// Overwrite this variable to change the client used.
var MainClient *http.Client

// TODO:
var BaseURL string

// WebClient is the interface used to represent a Client that performs requests.
// type WebClient interface {
// 	PerformGet(resource string) (*http.Response, error)
// 	PerformPost(resource string, body io.Reader) (*http.Response, error)
// }

// HttpClient is a Client that performs HTTP requests.
// type HttpClient struct {
// 	client  *http.Client
// 	baseurl string
// 	token   string
// }

// NewHttpClient creates a new HttpClient
// func NewHttpClient(baseurl string, timeoutSecs int, token string) WebClient {

// 	clientCredentialsConfig.Client(context.

// 	return HttpClient{
// 		client:  &http.Client{Timeout: time.Duration(timeoutSecs) * time.Second},
// 		baseurl: baseurl,
// 		token:   token,
// 	}
// }

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
