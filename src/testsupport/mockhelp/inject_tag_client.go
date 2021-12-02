package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/testsupport/mocks/sdkmocks"
)

const TAGID = "mock_id"

func PrepareTagMockClient(test_framework *testing.T) *sdkmocks.MockTagSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockTagSdkClient(ctrl)

	tags.Client = mockClient

	return mockClient.EXPECT()
}
