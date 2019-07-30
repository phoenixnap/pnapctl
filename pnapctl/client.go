package pnapctl

import (
	"io"
	"net/http"
	"time"
)

type WebClient interface {
	PerformGet(resource string) (*http.Response, error)
	PerformPost(resource string, body io.Reader) (*http.Response, error)
}

type HttpClient struct {
	client  *http.Client
	baseurl string
}

func NewHttpClient(baseurl string, timeout_secs int) WebClient {
	return HttpClient{
		client:  &http.Client{Timeout: time.Duration(timeout_secs) * time.Second},
		baseurl: baseurl,
	}
}

func (m HttpClient) PerformGet(resource string) (*http.Response, error) {
	return m.client.Get(m.buildURI(resource))
}

func (m HttpClient) PerformPost(resource string, body io.Reader) (*http.Response, error) {
	return m.client.Post(m.buildURI(resource), "application/json", body)
}

func (m HttpClient) buildURI(resource string) string {
	return m.baseurl + resource
}
