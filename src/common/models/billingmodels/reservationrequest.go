package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ReservationRequest struct{}

func p4() {
	x := billingapi.ReservationRequest{}
	fmt.Println("%v", x)
}
