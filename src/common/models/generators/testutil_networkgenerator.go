package generators

import (
	"phoenixnap.com/pnapctl/common/models/queryparams/network"
)

var GeneratePublicNetworksGetQueryParams = Generator(func(item *network.PublicNetworksGetQueryParams) {
	item.Location = &network.AllowedLocations[0]
})
