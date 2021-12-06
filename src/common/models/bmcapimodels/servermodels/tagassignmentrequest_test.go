package servermodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestTagAssignmentRequestToSdk(test_framework *testing.T) {
	tagAssignmentRequest := GenerateTagAssignmentRequestCli()
	sdkModel := tagAssignmentRequest.toSdk()

	assert.Equal(test_framework, tagAssignmentRequest.Name, sdkModel.Name)
	assert.Equal(test_framework, tagAssignmentRequest.Value, sdkModel.Value)
}
