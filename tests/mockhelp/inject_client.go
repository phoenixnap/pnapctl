package mockhelp

import (
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/mocks"
)

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
