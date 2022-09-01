package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationFromSdk(t *testing.T) {
	sdk := GenerateReservation()
	cli := ReservationFromSdk(sdk)

	assert.NotNil(t, sdk)
	assert.NotNil(t, cli)

	assertEqualReservation(t, sdk, cli)
}

func TestReservationFromSdk_Nil(t *testing.T) {
	assert.Nil(t, ReservationFromSdk(nil))
}

func assertEqualReservation(t *testing.T, sdk *billingapi.Reservation, cli *Reservation) {
	assert.Equal(t, sdk.Id, cli.Id)
	assert.Equal(t, sdk.ProductCode, cli.ProductCode)
	assert.Equal(t, sdk.ProductCategory, cli.ProductCategory)
	assert.Equal(t, sdk.Location, cli.Location)
	assert.Equal(t, sdk.ReservationModel, cli.ReservationModel)
	assert.Equal(t, sdk.InitialInvoiceModel, cli.InitialInvoiceModel)
	assert.Equal(t, sdk.StartDateTime, cli.StartDateTime)
	assert.Equal(t, sdk.EndDateTime, cli.EndDateTime)
	assert.Equal(t, sdk.LastRenewalDateTime, cli.LastRenewalDateTime)
	assert.Equal(t, sdk.NextRenewalDateTime, cli.NextRenewalDateTime)
	assert.Equal(t, sdk.AutoRenew, cli.AutoRenew)
	assert.Equal(t, sdk.Sku, cli.Sku)
	assert.Equal(t, sdk.Price, cli.Price)
	assert.Equal(t, sdk.PriceUnit, cli.PriceUnit)
	assert.Equal(t, sdk.AssignedResourceId, cli.AssignedResourceId)
}
