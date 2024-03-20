package tables

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestReservationTableFromSdk_NotNil(t *testing.T) {
	sdk := generators.Generate[billingapi.Reservation]()
	tbl := ReservationTableFromSdk(sdk)

	assertReservationTablesEqual(t, sdk, tbl)
}

func assertReservationTablesEqual(t *testing.T, sdk billingapi.Reservation, tbl ReservationTable) {
	assert.Equal(t, sdk.Id, tbl.Id)
	assert.Equal(t, sdk.ProductCode, tbl.ProductCode)
	assert.Equal(t, string(sdk.ProductCategory), tbl.ProductCategory)
	assert.Equal(t, string(sdk.Location), tbl.Location)
	assert.Equal(t, string(sdk.ReservationModel), tbl.ReservationModel)
	assert.Equal(t, string(*sdk.InitialInvoiceModel), tbl.InitialInvoiceModel)
	assert.Equal(t, sdk.StartDateTime.String(), tbl.StartDateTime)
	assert.Equal(t, sdk.EndDateTime.String(), tbl.EndDateTime)
	assert.Equal(t, sdk.LastRenewalDateTime.String(), tbl.LastRenewalDateTime)
	assert.Equal(t, sdk.NextRenewalDateTime.String(), tbl.NextRenewalDateTime)
	assert.Equal(t, sdk.AutoRenew, tbl.AutoRenew)
	assert.Equal(t, sdk.Sku, tbl.Sku)
	assert.Equal(t, sdk.Price, tbl.Price)
	assert.Equal(t, string(sdk.PriceUnit), tbl.PriceUnit)
	assert.Equal(t, *sdk.AssignedResourceId, tbl.AssignedResourceId)
}
