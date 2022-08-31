package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ReservationAutoRenewDisableRequest struct{}

func p3() {
	x := billingapi.ReservationAutoRenewDisableRequest{}
	fmt.Println("%v", x)
}
