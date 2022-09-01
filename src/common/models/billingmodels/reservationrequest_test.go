package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationRequestFromSdk_NotNil(t *testing.T) {
	sdk := GenerateReservationRequest()
	cli := ReservationRequestFromSdk(sdk)

	assert.NotNil(t, sdk)
	assert.NotNil(t, cli)

	assertEqualReservationRequestFromSdk(t, sdk, cli)
}

func TestReservationRequestFromSdk_Nil(t *testing.T) {
	assert.Nil(t, ReservationRequestFromSdk(nil))
}

func assertEqualReservationRequestFromSdk(t *testing.T, sdk *billingapi.ReservationRequest, cli *ReservationRequest) {
	assert.Equal(t, sdk.Sku, cli.Sku)
}
