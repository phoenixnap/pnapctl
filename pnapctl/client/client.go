package client

import (
	"context"
	"io"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnap-cli/pnapctl/configuration"
)

// WebClient is the interface used to represent a Client that performs requests.
type WebClient interface {
	PerformGet(resource string) (*http.Response, error)
	PerformPost(resource string, body io.Reader) (*http.Response, error)
}

// HTTPClient represents a struct containing an HTTP client and the baseURL
type HTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPClient creates a new instance of the HTTPClient struct using a client which performs the client credentials grant before a request
func NewHTTPClient(clientID string, clientSecret string) WebClient {
	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
	}

	httpClient := config.Client(context.Background())

	return HTTPClient{
		client:  httpClient,
		baseURL: configuration.Hostname,
	}
}

// MainClient is the main WebClient used to perform requests.
// Overwrite this variable to change the client used.
var MainClient WebClient

// PerformGet performs a Get request
func (m HTTPClient) PerformGet(resource string) (*http.Response, error) {
	return m.client.Get(m.buildURI(resource))
}

// PerformPost performs a Post request
func (m HTTPClient) PerformPost(resource string, body io.Reader) (*http.Response, error) {
	return m.client.Post(m.buildURI(resource), "application/json", body)
}

func (m HTTPClient) buildURI(resource string) string {
	return m.baseURL + resource
}
