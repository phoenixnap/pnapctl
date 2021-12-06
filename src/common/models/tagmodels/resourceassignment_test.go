package tagmodels

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
)

// tests
func TestResourceAssignmentFromSdk(test_framework *testing.T) {
	sdkResourceAssignment := GenerateResourceAssignmentSdk()
	resourceAssignment := ResourceAssignmentFromSdk(sdkResourceAssignment)

	assertEqualResourceAssignment(test_framework, *resourceAssignment, *sdkResourceAssignment)
}

func TestNilResourceAssignmentFromSdk(test_framework *testing.T) {
	resourceAssignment := ResourceAssignmentFromSdk(nil)

	assert.Nil(test_framework, resourceAssignment)
}

// assertion functions
func assertEqualResourceAssignment(test_framework *testing.T, r1 ResourceAssignment, r2 tagapisdk.ResourceAssignment) {
	assert.Equal(test_framework, r1.ResourceName, r2.ResourceName)
	assert.Equal(test_framework, r1.Value, r2.Value)
}
