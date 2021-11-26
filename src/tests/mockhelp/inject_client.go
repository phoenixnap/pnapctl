package mockhelp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func WithResponse(status int, body io.ReadCloser) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       body,
	}
}

func WithBody(body interface{}) io.ReadCloser {
	data, _ := json.Marshal(body)

	return ioutil.NopCloser(bytes.NewBuffer(data))
}
