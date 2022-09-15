package networkmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

var allowedLocations = []string{
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
	if location != "" && !iterutils.Contains(allowedLocations, location) {
		return nil, ctlerrors.InvalidFlagValuePassedError("location", location, allowedLocations)
	}

	var validLocation *string
	if location != "" {
		if !iterutils.Contains(allowedLocations, location) {
			return nil, ctlerrors.InvalidFlagValuePassedError("location", location, allowedLocations)
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
