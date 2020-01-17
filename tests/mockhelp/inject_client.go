package mockhelp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/common/client"
	"phoenixnap.com/pnap-cli/tests/mocks"
)

var Body io.Writer
var URL string

const SERVERID = "mock_id"

func PrepareMockClient(test_framework *testing.T) *mocks.MockWebClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := mocks.NewMockWebClient(ctrl)

	client.MainClient = mockClient

	return mockClient.EXPECT()
}

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
