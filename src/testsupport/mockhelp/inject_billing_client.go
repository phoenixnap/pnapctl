package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/testsupport/mocks/sdkmocks"
)

func PrepareBillingMockClient(test_framework *testing.T) *sdkmocks.MockBillingSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockBillingSdkClient(ctrl)

	billing.Client = mockClient

	return mockClient.EXPECT()
}
