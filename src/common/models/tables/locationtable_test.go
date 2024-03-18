package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/phoenixnap/go-sdk-bmc/locationapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestToLocationTable(test_framework *testing.T) {
	location := generators.Generate[locationapi.Location]()
	table := ToLocationTable(location)

	assertLocationsEqual(test_framework, location, table)
}

func assertLocationsEqual(test_framework *testing.T, location locationapi.Location, table Location) {
	assert.Equal(test_framework, string(location.Location), table.Location)
	assert.Equal(test_framework, *location.LocationDescription, table.LocationDescription)
	assert.Equal(test_framework, iterutils.MapRef(location.ProductCategories, models.ProductCategoryToTableString), table.ProductCategories)
}
