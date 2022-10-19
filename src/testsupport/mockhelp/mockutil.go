package mockhelp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const FILENAME = "testfile.yaml"
const RESOURCEID = "mock_id"

func WithResponse(status int, body io.ReadCloser) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       body,
	}
}

func WithBody(body interface{}) io.ReadCloser {
	data, _ := json.Marshal(body)

	return io.NopCloser(bytes.NewBuffer(data))
}
