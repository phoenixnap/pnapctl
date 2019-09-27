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

type HttpClient struct {
	client  *http.Client
	baseurl string
}

func NewHttpClient(clientId string, clientSecret string) WebClient {
	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
	}

	httpClient := config.Client(context.Background())

	return HttpClient{
		client:  httpClient,
		baseurl: configuration.Hostname,
	}
}

// MainClient is the main WebClient used to perform requests.
// Overwrite this variable to change the client used.
var MainClient WebClient

// PerformGet performs a Get request
func (m HttpClient) PerformGet(resource string) (*http.Response, error) {
	return m.client.Get(m.buildURI(m.baseurl, resource))
}

// PerformPost performs a Post request
func (m HttpClient) PerformPost(resource string, body io.Reader) (*http.Response, error) {
	return m.client.Post(m.buildURI(m.baseurl, resource), "application/json", body)
}

func (m HttpClient) buildURI(baseURL string, resource string) string {
	return baseURL + resource
}
