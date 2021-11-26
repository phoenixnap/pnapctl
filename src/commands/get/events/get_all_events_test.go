package events

import (
	"errors"
	"testing"
	"time"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/auditmodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllEventsSuccess(test_framework *testing.T) {
	eventList := generators.GenerateEvents(2)
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)

	var eventTables []interface{}

	for _, event := range eventList {
		eventTables = append(eventTables, tables.ToEventTable(event))
	}

	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet(queryParams).
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
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet(queryParams).
		Return(event, nil, testutil.TestKeycloakError)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllEventsPrinterFailure(test_framework *testing.T) {
	eventList := generators.GenerateEvents(2)
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)

	var eventTables []interface{}

	for _, event := range eventList {
		eventTables = append(eventTables, tables.ToEventTable(event))
	}

	PrepareAuditMockClient(test_framework).
		EventsGet(queryParams).
		Return(eventList, WithResponse(200, WithBody(eventList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables, "get events").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetEventsServerError(test_framework *testing.T) {
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)

	PrepareAuditMockClient(test_framework).
		EventsGet(queryParams).
		Return([]auditapisdk.Event{}, WithResponse(500, nil), nil)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get events' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetEventsInvalidTimeFormat(test_framework *testing.T) {
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)
	From = "Not A Date"

	err := GetEventsCmd.RunE(GetEventsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "'From' (Not A Date) is not a valid date."
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetEventsInvalidOrderFormat(test_framework *testing.T) {
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)
	Order = "None"

	err := GetEventsCmd.RunE(GetEventsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Invalid Order 'None'. Valid values: 'ASC', 'DESC'"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetEventsInvalidVerbFormat(test_framework *testing.T) {
	queryParams := generators.GenerateQueryParams()
	setQueryParams(queryParams)
	Verb = "Doing"

	err := GetEventsCmd.RunE(GetEventsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Invalid Verb 'Doing'. Valid values: 'POST', 'PUT', 'PATCH', 'DELETE'"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func setQueryParams(queryparams auditmodels.EventsGetQueryParams) {
	From = queryparams.From.Format(time.RFC3339)
	To = queryparams.To.Format(time.RFC3339)
	Limit = queryparams.Limit
	Order = queryparams.Order
	Username = queryparams.Username
	Verb = queryparams.Verb
	Uri = queryparams.Uri
}
