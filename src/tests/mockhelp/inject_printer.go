package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/tests/mocks"
)

func PrepareMockPrinter(test_framework *testing.T) *mocks.MockPrinterMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockPrinter := mocks.NewMockPrinter(ctrl)

	printer.MainPrinter = mockPrinter

	return mockPrinter.EXPECT()
}
