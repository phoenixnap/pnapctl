package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ReservationAutoRenewDisableRequest struct {
	AutoRenewDisableReason *string `json:"autoRenewDisableReason,omitempty" yaml:"autoRenewDisableReason,omitempty"`
}

func ReservationAutoRenewDisableRequestFromSdk(sdk billingapi.ReservationAutoRenewDisableRequest) ReservationAutoRenewDisableRequest {
	return ReservationAutoRenewDisableRequest{
		AutoRenewDisableReason: sdk.AutoRenewDisableReason,
	}
}

func (r *ReservationAutoRenewDisableRequest) ToSdk() billingapi.ReservationAutoRenewDisableRequest {
	return billingapi.ReservationAutoRenewDisableRequest{
		AutoRenewDisableReason: r.AutoRenewDisableReason,
	}
}

func CreateReservationAutoRenewDisableRequestFromFile(filename string, commandname string) (*billingapi.ReservationAutoRenewDisableRequest, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	var reservationAutoRenewDisableRequest ReservationAutoRenewDisableRequest

	err = files.Unmarshal(data, &reservationAutoRenewDisableRequest, commandname)

	if err != nil {
		return nil, err
	}

	sdkAutoRenewDisableRequest := reservationAutoRenewDisableRequest.ToSdk()

	return &sdkAutoRenewDisableRequest, nil
}
