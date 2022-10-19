package mockhelp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"phoenixnap.com/pnapctl/testsupport/testutil"
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

// File Processor Mocks
func ExpectFromFileSuccess(t *testing.T, marshaller func(interface{}) ([]byte, error), item interface{}) {
	marshalled, _ := marshaller(item)

	PrepareMockFileProcessor(t).
		ReadFile(FILENAME).
		Return(marshalled, nil)
}

func ExpectFromFileFailure(t *testing.T) error {
	PrepareMockFileProcessor(t).
		ReadFile(FILENAME).
		Return(nil, testutil.TestError)

	return testutil.TestError
}

func ExpectFromFileUnmarshalFailure(t *testing.T) {
	PrepareMockFileProcessor(t).
		ReadFile(FILENAME).
		Return([]byte(`Invalid JSON/YAML - Should cause unmarshal to fail.`), nil)
}

// Printer Mocks
func ExpectToPrintSuccess(t *testing.T, item interface{}) {
	PrepareMockPrinter(t).
		PrintOutput(item).
		Return(nil)
}

func ExpectToPrintFailure(t *testing.T, item interface{}) error {
	PrepareMockPrinter(t).
		PrintOutput(item).
		Return(testutil.TestError)

	return testutil.TestError
}
