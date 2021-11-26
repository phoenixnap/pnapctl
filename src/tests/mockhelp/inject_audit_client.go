package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/common/client/audit"
	"phoenixnap.com/pnap-cli/tests/mocks/sdkmocks"
)

func PrepareAuditMockClient(test_framework *testing.T) *sdkmocks.MockAuditSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockAuditSdkClient(ctrl)

	audit.Client = mockClient

	return mockClient.EXPECT()
}
