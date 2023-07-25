package mockhelp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnapctl/common/client/audit"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/client/locations"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/client/networkstorage"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/fileprocessor"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/testsupport/mocks"
	"phoenixnap.com/pnapctl/testsupport/mocks/sdkmocks"
)

func PrepareMockFileProcessor(test_framework *testing.T) *mocks.MockFileProcessorMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockFileProcessor := mocks.NewMockFileProcessor(ctrl)

	fileprocessor.MainFileProcessor = mockFileProcessor

	return mockFileProcessor.EXPECT()
}

func PrepareMockPrinter(test_framework *testing.T) *mocks.MockPrinterMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockPrinter := mocks.NewMockPrinter(ctrl)

	printer.MainPrinter = mockPrinter

	return mockPrinter.EXPECT()
}

func PrepareAuditMockClient(test_framework *testing.T) *sdkmocks.MockAuditSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockAuditSdkClient(ctrl)

	audit.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareBillingMockClient(test_framework *testing.T) *sdkmocks.MockBillingSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockBillingSdkClient(ctrl)

	billing.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareBmcApiMockClient(test_framework *testing.T) *sdkmocks.MockBmcApiSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockBmcApiSdkClient(ctrl)

	bmcapi.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareIPMockClient(test_framework *testing.T) *sdkmocks.MockIpSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockIpSdkClient(ctrl)

	ip.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareNetworkMockClient(test_framework *testing.T) *sdkmocks.MockNetworkSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockNetworkSdkClient(ctrl)

	networks.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareNetworkStorageApiMockClient(test_framework *testing.T) *sdkmocks.MockNetworkStorageSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockNetworkStorageSdkClient(ctrl)

	networkstorage.Client = mockClient

	return mockClient.EXPECT()
}
func PrepareRancherMockClient(test_framework *testing.T) *sdkmocks.MockRancherSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockRancherSdkClient(ctrl)

	rancher.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareTagMockClient(test_framework *testing.T) *sdkmocks.MockTagSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockTagSdkClient(ctrl)

	tags.Client = mockClient

	return mockClient.EXPECT()
}

func PrepareLocationMockClient(test_framework *testing.T) *sdkmocks.MockLocationSdkClientMockRecorder {
	ctrl := gomock.NewController(test_framework)
	mockClient := sdkmocks.NewMockLocationSdkClient(ctrl)

	locations.Client = mockClient

	return mockClient.EXPECT()
}
