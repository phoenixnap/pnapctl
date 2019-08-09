package client

import (
	"io"
	"net/http"
	"time"
)

// MainClient is the main WebClient used to perform requests.
// Overwrite this variable to change the client used.
var MainClient = NewHttpClient("https://localhost:8080", 10)

// WebClient is the interface used to represent a Client that performs requests.
type WebClient interface {
	PerformGet(resource string) (*http.Response, error)
	PerformPost(resource string, body io.Reader) (*http.Response, error)
}

// HttpClient is a Client that performs HTTP requests.
type HttpClient struct {
	client  *http.Client
	baseurl string
}

// NewHttpClient creates a new HttpClient
func NewHttpClient(baseurl string, timeoutSecs int) WebClient {
	return HttpClient{
		client:  &http.Client{Timeout: time.Duration(timeoutSecs) * time.Second},
		baseurl: baseurl,
	}
}

// PerformGet performs a Get request
func (m HttpClient) PerformGet(resource string) (*http.Response, error) {
	return m.client.Get(m.buildURI(resource))
}

// PerformPost performs a Post request
func (m HttpClient) PerformPost(resource string, body io.Reader) (*http.Response, error) {
	return m.client.Post(m.buildURI(resource), "application/json", body)
}

func (m HttpClient) buildURI(resource string) string {
	return m.baseurl + resource
}
