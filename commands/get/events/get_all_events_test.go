package events

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	auditapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/auditapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllEventsSuccess(test_framework *testing.T) {
	eventList := generators.GenerateEvents(2)

	var eventTables []interface{}

	for _, event := range eventList {
		eventTables = append(eventTables, tables.ToEventTable(event))
	}

	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet().
		Return(eventList, WithResponse(200, WithBody(eventList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables, "get events").
		Return(nil)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllEventsKeycloakFailure(test_framework *testing.T) {
	event := []auditapisdk.Event{generators.GenerateEvent()}
	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet().
		Return(event, nil, testutil.TestKeycloakError)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllEventsPrinterFailure(test_framework *testing.T) {
	eventList := generators.GenerateEvents(2)

	var eventTables []interface{}

	for _, event := range eventList {
		eventTables = append(eventTables, tables.ToEventTable(event))
	}

	PrepareAuditMockClient(test_framework).
		EventsGet().
		Return(eventList, WithResponse(200, WithBody(eventList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables, "get events").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetEventsServerError(test_framework *testing.T) {
	PrepareAuditMockClient(test_framework).
		EventsGet().
		Return([]auditapisdk.Event{}, WithResponse(500, nil), nil)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get events' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}
