package mockhelp

import (
	"phoenixnap.com/pnapctl/common/client/ip"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/testsupport/mocks/sdkmocks"
)

const IPBLOCKID = "mock_id"

func PrepareIPMockClient(test_framework *testing.T) *sdkmocks.MockIpSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockIpSdkClient(ctrl)

	ip.Client = mockClient

	return mockClient.EXPECT()
}
