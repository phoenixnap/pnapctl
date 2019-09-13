package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/pnapctl/fileprocessor"

	"phoenixnap.com/pnap-cli/pnapctl/mocks"
)

const FILENAME = "somefile.yaml"

func PrepareMockFileProcessor(test_framework *testing.T) *mocks.MockFileProcessorMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockFileProcessor := mocks.NewMockFileProcessor(ctrl)

	fileprocessor.MainFileProcessor = mockFileProcessor

	return mockFileProcessor.EXPECT()
}
