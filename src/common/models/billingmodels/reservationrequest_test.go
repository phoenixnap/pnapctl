package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationRequestFromSdk_NotNil(t *testing.T) {
	cli := *GenerateReservationRequestCli()
	sdk := cli.ToSdk()

	assertEqualReservationRequestFromSdk(t, cli, sdk)
}

func assertEqualReservationRequestFromSdk(t *testing.T, cli ReservationRequest, sdk billingapi.ReservationRequest) {
	assert.Equal(t, cli.Sku, sdk.Sku)
}
