package events

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getRequestParams() (string, string, int, string, string, string, string) {
	return From, To, Limit, Order, Username, Verb, Uri
}

func TestGetAllEventsSuccess(test_framework *testing.T) {
	eventList := testutil.GenN(2, generators.Generate[auditapi.Event])

	var eventTables []interface{}

	for _, event := range eventList {
		eventTables = append(eventTables, tables.ToEventTable(event))
	}

	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(eventList, WithResponse(200, WithBody(eventList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables).
		Return(nil)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllEventsKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllEventsPrinterFailure(test_framework *testing.T) {
	eventList := testutil.GenN(2, generators.Generate[auditapi.Event])

	var eventTables []interface{}

	for _, event := range eventList {
		eventTables = append(eventTables, tables.ToEventTable(event))
	}

	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(eventList, WithResponse(200, WithBody(eventList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetEventsServerError(test_framework *testing.T) {
	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(nil, WithResponse(500, nil), nil)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}
