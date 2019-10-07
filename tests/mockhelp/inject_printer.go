package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/pnapctl/mocks"
	"phoenixnap.com/pnap-cli/pnapctl/printer"
)

func PrepareMockPrinter(test_framework *testing.T) *mocks.MockPrinterMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockPrinter := mocks.NewMockPrinter(ctrl)

	printer.MainPrinter = mockPrinter

	return mockPrinter.EXPECT()
}
