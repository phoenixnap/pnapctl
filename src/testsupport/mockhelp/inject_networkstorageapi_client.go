package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/testsupport/mocks/sdkmocks"
)

func PrepareNetworkStorageApiMockClient(test_framework *testing.T) *sdkmocks.MockNetworkStorageSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockNetworkStorageSdkClient(ctrl)

	networkstorage.Client = mockClient

	return mockClient.EXPECT()
}
