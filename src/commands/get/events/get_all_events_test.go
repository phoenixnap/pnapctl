package events

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getRequestParams() (string, string, int, string, string, string, string) {
	return From, To, Limit, Order, Username, Verb, Uri
}

func TestGetAllEventsSuccess(test_framework *testing.T) {
	eventList := testutil.GenN(2, generators.Generate[auditapi.Event])
	eventTables := iterutils.MapInterface(eventList, tables.ToEventTable)

	// Mocking
	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(eventList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables).
		Return(nil)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllEventsClientFailure(test_framework *testing.T) {
	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(nil, testutil.TestError)

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllEventsPrinterFailure(test_framework *testing.T) {
	eventList := testutil.GenN(2, generators.Generate[auditapi.Event])
	eventTables := iterutils.MapInterface(eventList, tables.ToEventTable)

	PrepareAuditMockClient(test_framework).
		EventsGet(getRequestParams()).
		Return(eventList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(eventTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetEventsCmd.RunE(GetEventsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
