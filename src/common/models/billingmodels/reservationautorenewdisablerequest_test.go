package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationAutoRenewDisableRequestFromSdk_NotNil(t *testing.T) {
	sdk := GenerateReservationAutoRenewDisableRequest()
	cli := ReservationAutoRenewDisableRequestFromSdk(sdk)

	assert.NotNil(t, sdk)
	assert.NotNil(t, cli)

	assertEqualReservationAutoRenewDisableRequestFromSdk(t, sdk, cli)
}

func TestReservationAutoRenewDisableRequestFromSdk_Nil(t *testing.T) {
	assert.Nil(t, ReservationAutoRenewDisableRequestFromSdk(nil))
}

func assertEqualReservationAutoRenewDisableRequestFromSdk(t *testing.T, sdk *billingapi.ReservationAutoRenewDisableRequest, cli *ReservationAutoRenewDisableRequest) {
	assert.Equal(t, sdk.AutoRenewDisableReason, cli.AutoRenewDisableReason)
}
