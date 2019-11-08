package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnap-cli/pnapctl/configuration"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
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

// HandleClientResponse determines whether
// (i) request was executed at all, if not return a generic error
// (ii) request was executed sucesfully, if so unmarshall response, print it and don't return an error
// (iii) request was executed unsucesfully, if so attempt to use the return body to output an error
func HandleClientResponse(response *http.Response, err error, commandName string) error {
	if response == nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.IncorrectRequestStructure)
	} else if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return ctlerrors.GenericNonRequestError(ctlerrors.ResponseBodyReadFailure, commandName)
		}

		responseBody := ResponseBody{}
		err = json.Unmarshal(body, &responseBody)

		if err != nil {
			return ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingResponseBody, commandName)
		}

		fmt.Println(responseBody.Result)
		return nil
	} else {
		return ctlerrors.HandleResponseError(response, commandName)
	}
}

// executeRequest will perform the http request provided and return the result
// with the error decorated accordingly IF it is an auth error
func executeRequest(f func() (*http.Response, error)) (*http.Response, error) {
	response, err := f()

	if e, isUrlError := err.(*url.Error); isUrlError && strings.Contains(err.Error(), "oauth2: cannot fetch token") {
		//Timeout If there is an error it must have happened while resolving token
		// ErrorURLs frome the actual request should be represented in the body
		return response, ctlerrors.Error{Msg: "Failed to resolve provided credentials", Cause: e}
	}

	return response, err
}

func (m HTTPClient) buildURI(resource string) string {
	return m.baseURL + resource
}
