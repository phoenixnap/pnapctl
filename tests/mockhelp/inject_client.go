package mockhelp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/tests/mocks/sdkmocks"
)

var Body io.Writer

const RESOURCEID = "mock_id"

func PrepareBmcApiMockClient(test_framework *testing.T) *sdkmocks.MockBmcApiSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockBmcApiSdkClient(ctrl)

	bmcapi.Client = mockClient

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
