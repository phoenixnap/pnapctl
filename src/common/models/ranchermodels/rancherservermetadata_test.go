package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestRancherServerMetadataToSdk(test_framework *testing.T) {
	rancherServerMetadata := GenerateRancherServerMetadataCli()
	sdkRancherServerMetadata := *rancherServerMetadata.ToSdk()

	assertEqualRancherServerMetadata(test_framework, rancherServerMetadata, sdkRancherServerMetadata)
}

func TestRancherServerMetadataFromSdk(test_framework *testing.T) {
	sdkRancherServerMetadata := GenerateRancherServerMetadataSdk()
	rancherServerMetadata := *RancherServerMetadataFromSdk(&sdkRancherServerMetadata)

	assertEqualRancherServerMetadata(test_framework, rancherServerMetadata, sdkRancherServerMetadata)
}

func TestRancherServerMetadataToTableString_nilMetadata(test_framework *testing.T) {
	result := RancherServerMetadataToTableString(nil)
	assert.Equal(test_framework, "", result)
}

func TestRancherServerMetadataToTableString_urlOnly(test_framework *testing.T) {
	sdkModel := ranchersdk.RancherServerMetadata{
		Url:      testutil.RandSeqPointer(10),
		Password: nil,
		Username: nil,
	}
	result := RancherServerMetadataToTableString(&sdkModel)
	assert.Equal(test_framework, "Url: "+*sdkModel.Url+"\n", result)
}

func TestRancherServerMetadataToTableString_passwordOnly(test_framework *testing.T) {
	sdkModel := ranchersdk.RancherServerMetadata{
		Url:      nil,
		Password: testutil.RandSeqPointer(10),
		Username: nil,
	}
	result := RancherServerMetadataToTableString(&sdkModel)
	assert.Equal(test_framework, "Pass: "+*sdkModel.Password+"\n", result)
}

func TestRancherServerMetadataToTableString_usernameOnly(test_framework *testing.T) {
	sdkModel := ranchersdk.RancherServerMetadata{
		Url:      nil,
		Password: nil,
		Username: testutil.RandSeqPointer(10),
	}
	result := RancherServerMetadataToTableString(&sdkModel)
	assert.Equal(test_framework, "User: "+*sdkModel.Username+"\n", result)
}

func TestRancherServerMetadataToTableString_fullMetadata(test_framework *testing.T) {
	sdkModel := ranchersdk.RancherServerMetadata{
		Url:      testutil.RandSeqPointer(10),
		Password: testutil.RandSeqPointer(10),
		Username: testutil.RandSeqPointer(10),
	}
	result := RancherServerMetadataToTableString(&sdkModel)

	expectedResult := "User: " + *sdkModel.Username + "\nPass: " + *sdkModel.Password + "\nUrl: " + *sdkModel.Url + "\n"

	assert.Equal(test_framework, expectedResult, result)
}

func assertEqualRancherServerMetadata(test_framework *testing.T, r1 RancherServerMetadata, r2 ranchersdk.RancherServerMetadata) {
	assert.Equal(test_framework, r1.Url, r2.Url)
	assert.Equal(test_framework, r1.Username, r2.Username)
	assert.Equal(test_framework, r1.Password, r2.Password)
}
