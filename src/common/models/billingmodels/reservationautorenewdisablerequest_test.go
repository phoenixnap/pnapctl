package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
)

func TestReservationAutoRenewDisableRequestToSdk_NotNil(t *testing.T) {
	cli := *GenerateReservationAutoRenewDisableRequestCli()
	sdk := cli.ToSdk()

	assertEqualReservationAutoRenewDisableRequestFromSdk(t, cli, sdk)
}

func assertEqualReservationAutoRenewDisableRequestFromSdk(t *testing.T, cli ReservationAutoRenewDisableRequest, sdk billingapi.ReservationAutoRenewDisableRequest) {
	assert.Equal(t, cli.AutoRenewDisableReason, sdk.AutoRenewDisableReason)
}
