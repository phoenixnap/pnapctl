package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/tests/mocks/sdkmocks"
)

const RESOURCEID = "mock_id"

func PrepareBmcApiMockClient(test_framework *testing.T) *sdkmocks.MockBmcApiSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockBmcApiSdkClient(ctrl)

	bmcapi.Client = mockClient

	return mockClient.EXPECT()
}
