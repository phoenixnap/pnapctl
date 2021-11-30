package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/tests/mocks/sdkmocks"
)

const NETWORKID = "mock_id"

func PrepareNetworkMockClient(test_framework *testing.T) *sdkmocks.MockNetworkSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockNetworkSdkClient(ctrl)

	networks.Client = mockClient

	return mockClient.EXPECT()
}
