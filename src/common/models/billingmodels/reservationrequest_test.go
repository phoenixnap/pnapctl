package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationRequestFromSdk_NotNil(t *testing.T) {
	sdk := *GenerateReservationRequest()
	cli := ReservationRequestFromSdk(sdk)

	assertEqualReservationRequestFromSdk(t, sdk, cli)
}

func assertEqualReservationRequestFromSdk(t *testing.T, sdk billingapi.ReservationRequest, cli ReservationRequest) {
	assert.Equal(t, sdk.Sku, cli.Sku)
}
