package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/testsupport/mocks/sdkmocks"
)

const CLUSTERID = "mock_id"

func PrepareRancherMockClient(test_framework *testing.T) *sdkmocks.MockRancherSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockRancherSdkClient(ctrl)

	rancher.Client = mockClient

	return mockClient.EXPECT()
}
