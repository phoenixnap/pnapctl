package auditmodels

import (
	"testing"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/assert"
)

// tests
func TestEventFromSdk(test_framework *testing.T) {
	sdkEvent := GenerateEventSdk()
	event := EventFromSdk(sdkEvent)

	assertEqualEvent(test_framework, *event, *sdkEvent)
}

func TestNilEventFromSdk(test_framework *testing.T) {
	event := EventFromSdk(nil)

	assert.Nil(test_framework, event)
}

// assertion functions
func assertEqualEvent(test_framework *testing.T, event Event, sdkEvent auditapisdk.Event) {
	assert.Equal(test_framework, event.Name, sdkEvent.Name)
	assert.Equal(test_framework, event.Timestamp, sdkEvent.Timestamp)

	assertEqualUserInfo(test_framework, event.UserInfo, sdkEvent.UserInfo)
}
