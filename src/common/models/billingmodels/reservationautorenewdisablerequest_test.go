package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationAutoRenewDisableRequestFromSdk_NotNil(t *testing.T) {
	sdk := *GenerateReservationAutoRenewDisableRequest()
	cli := ReservationAutoRenewDisableRequestFromSdk(sdk)

	assertEqualReservationAutoRenewDisableRequestFromSdk(t, sdk, cli)
}

func assertEqualReservationAutoRenewDisableRequestFromSdk(t *testing.T, sdk billingapi.ReservationAutoRenewDisableRequest, cli ReservationAutoRenewDisableRequest) {
	assert.Equal(t, sdk.AutoRenewDisableReason, cli.AutoRenewDisableReason)
}
