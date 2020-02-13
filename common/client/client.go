package client

import (
	"context"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	configuration "phoenixnap.com/pnap-cli/configs"
)

// MainClient is the main WebClient used to perform requests.
// Overwrite this variable to change the client used.
var MainClient WebClient

// WebClient is the interface used to represent a Client that performs requests.
type WebClient interface {
	PerformGet(resource string) (*http.Response, error)
	PerformPost(resource string, body io.Reader) (*http.Response, error)
	PerformDelete(resource string) (*http.Response, error)
}

// HTTPClient is a Client that performs HTTP requests.
type HTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPClient creates a new HTTPClient
func NewHTTPClient(clientID string, clientSecret string) WebClient {
	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	httpClient := config.Client(context.Background())

	return HTTPClient{
		client:  httpClient,
		baseURL: configuration.Hostname,
	}
}

// PerformGet performs a Get request and check for auth errors
func (m HTTPClient) PerformGet(resource string) (*http.Response, error) {
	return executeRequest(func() (*http.Response, error) {
		return m.client.Get(m.buildURI(resource))
	})
}

// PerformDelete performs a Delete request and check for auth errors
func (m HTTPClient) PerformDelete(resource string) (*http.Response, error) {
	return executeRequest(func() (*http.Response, error) {
		req, err := http.NewRequest("DELETE", m.buildURI(resource), nil)
		// replicating Get/Post error handling
		if err != nil {
			return nil, err
		}
		return m.client.Do(req)
	})
}

// PerformPost performs a Post request and check for auth errors
func (m HTTPClient) PerformPost(resource string, body io.Reader) (*http.Response, error) {
	return executeRequest(func() (*http.Response, error) {
		return m.client.Post(m.buildURI(resource), "application/json", body)
	})
}

// ResponseBody represents the format of the expected body returned from the server. it may have additional fields but we only care about the Result
type ResponseBody struct {
	Result string
}

// executeRequest will perform the http request provided and return the result
// with the error decorated accordingly IF it is an auth error
func executeRequest(f func() (*http.Response, error)) (*http.Response, error) {
	response, err := f()

	if err == nil || !os.IsTimeout(err) {
		return response, err
	} else {
		//Timeout If there is an error it must have happened while resolving token
		// ErrorURLs frome the actual request should be represented in the body
		return response, ctlerrors.CLIError{Message: "Failed to resolve provided credentials", Cause: err}
	}
}

func (m HTTPClient) buildURI(resource string) string {
	return m.baseURL + resource
}
