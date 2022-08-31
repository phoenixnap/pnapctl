package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ConfigurationDetails struct{}

func p2() {
	x := billingapi.ConfigurationDetails{}
	fmt.Println("%v", x)
}
