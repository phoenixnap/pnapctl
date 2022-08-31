package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ProductAvailability struct{}

func p1() {
	x := billingapi.ProductAvailability{}
	fmt.Println("%v", x)
}
