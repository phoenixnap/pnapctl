package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ReservationRequest struct {
	Sku string `json:"sku" yaml:"sku"`
}

func ReservationRequestFromSdk(sdk billingapi.ReservationRequest) ReservationRequest {
	return ReservationRequest{
		Sku: sdk.Sku,
	}
}

func (r *ReservationRequest) ToSdk() billingapi.ReservationRequest {
	return billingapi.ReservationRequest{
		Sku: r.Sku,
	}
}

func CreateReservationRequestFromFile(filename string, commandname string) (*billingapi.ReservationRequest, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	var reservationRequest ReservationRequest

	err = files.Unmarshal(data, &reservationRequest, commandname)

	if err != nil {
		return nil, err
	}

	sdkReservationRequest := reservationRequest.ToSdk()

	return &sdkReservationRequest, nil
}
