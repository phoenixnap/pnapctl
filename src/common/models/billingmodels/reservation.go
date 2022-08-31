package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type Reservation struct{}

func p5() {
	x := billingapi.ReservationRequest{}
	fmt.Println("%v", x)
}
