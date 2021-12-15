package servermodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func TestTagAssignmentRequestToSdk(test_framework *testing.T) {
	tagAssignmentRequest := GenerateTagAssignmentRequestCli()
	sdkModel := tagAssignmentRequest.toSdk()

	assertEqualTagAssignmentRequest(test_framework, tagAssignmentRequest, sdkModel)

}

func assertEqualTagAssignmentRequest(test_framework *testing.T, cliTagAssignmentRequest TagAssignmentRequest, sdkTagAssignmentRequest bmcapisdk.TagAssignmentRequest) {
	assert.Equal(test_framework, cliTagAssignmentRequest.Name, sdkTagAssignmentRequest.Name)
	assert.Equal(test_framework, cliTagAssignmentRequest.Value, sdkTagAssignmentRequest.Value)

}
