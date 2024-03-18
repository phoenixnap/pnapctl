package locations

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/locationapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllLocationsSuccess(test_framework *testing.T) {
	locationList := testutil.GenN(2, generators.Generate[locationapi.Location])
	locationTables := iterutils.MapInterface(locationList, tables.ToLocationTable)

	// Mocking
	PrepareLocationMockClient(test_framework).
		LocationsGet(Location, ProductCategory).
		Return(locationList, nil)

	ExpectToPrintSuccess(test_framework, locationTables)

	err := GetLocationsCmd.RunE(GetLocationsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllLocationsClientFailure(test_framework *testing.T) {
	PrepareLocationMockClient(test_framework).
		LocationsGet(Location, ProductCategory).
		Return(nil, testutil.TestError)

	err := GetLocationsCmd.RunE(GetLocationsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllLocationsPrinterFailure(test_framework *testing.T) {
	locationList := testutil.GenN(2, generators.Generate[locationapi.Location])
	locationTables := iterutils.MapInterface(locationList, tables.ToLocationTable)

	// Mocking
	PrepareLocationMockClient(test_framework).
		LocationsGet(Location, ProductCategory).
		Return(locationList, nil)

	expectedErr := ExpectToPrintFailure(test_framework, locationTables)

	err := GetLocationsCmd.RunE(GetLocationsCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
