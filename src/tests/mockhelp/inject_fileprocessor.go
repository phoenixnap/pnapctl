package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/fileprocessor"

	"phoenixnap.com/pnapctl/tests/mocks"
)

const FILENAME = "testfile.yaml"

func PrepareMockFileProcessor(test_framework *testing.T) *mocks.MockFileProcessorMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockFileProcessor := mocks.NewMockFileProcessor(ctrl)

	fileprocessor.MainFileProcessor = mockFileProcessor

	return mockFileProcessor.EXPECT()
}
