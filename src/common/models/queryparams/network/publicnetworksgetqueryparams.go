package network

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

var AllowedLocations = []string{
	"PHX",
	"ASH",
	"SGP",
	"NLD",
	"CHI",
	"SEA",
	"AUS",
}

type PublicNetworksGetQueryParams struct {
	Location *string
}

func NewPublicNetworksGetQueryParams(location string) (*PublicNetworksGetQueryParams, error) {
	if location != "" && !iterutils.Contains(AllowedLocations, location) {
		return nil, ctlerrors.InvalidFlagValuePassedError("location", location, AllowedLocations)
	}

	var validLocation *string
	if location != "" {
		if !iterutils.Contains(AllowedLocations, location) {
			return nil, ctlerrors.InvalidFlagValuePassedError("location", location, AllowedLocations)
		}
		validLocation = &location
	}

	return &PublicNetworksGetQueryParams{
		Location: validLocation,
	}, nil
}

func (queryParams *PublicNetworksGetQueryParams) AttachToRequest(request networkapi.ApiPublicNetworksGetRequest) networkapi.ApiPublicNetworksGetRequest {
	if queryParams.Location != nil {
		request = request.Location(*queryParams.Location)
	}
	return request
}
