package billingmodels

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type ServerDetails struct {
	Id       string `json:"id" yaml:"id"`
	Hostname string `json:"hostname" yaml:"hostname"`
}

func ServerDetailsFromSdk(serverDetails *billingapi.ServerDetails) *ServerDetails {
	return &ServerDetails{
		Id:       serverDetails.Id,
		Hostname: serverDetails.Hostname,
	}
}
