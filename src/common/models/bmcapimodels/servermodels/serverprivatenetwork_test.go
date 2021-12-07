package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests

// TODO

// assertion functions
func assertEqualServerPrivateNetwork(test_framework *testing.T, cliServerPrivateNetwork ServerPrivateNetwork, sdkServerPrivateNetwork bmcapisdk.ServerPrivateNetwork) {
	assert.Equal(test_framework, cliServerPrivateNetwork.Id, sdkServerPrivateNetwork.Id)

	if cliServerPrivateNetwork.Ips == nil {
		assert.Nil(test_framework, sdkServerPrivateNetwork.Ips, "CLI Server Private Network's IPs are nil, but not SDK Server Private Network's IPs.")
	} else if sdkServerPrivateNetwork.Ips == nil {
		assert.Nil(test_framework, cliServerPrivateNetwork.Ips, "SDK Server Private Network's IPs are nil, but not Server Private Network's IPs.")
	} else {
		assert.Equal(test_framework, len(*cliServerPrivateNetwork.Ips), len(*sdkServerPrivateNetwork.Ips))

		for i := range *cliServerPrivateNetwork.Ips {
			assert.Equal(test_framework, (*cliServerPrivateNetwork.Ips)[i], (*sdkServerPrivateNetwork.Ips)[i])
		}
	}
	assert.Equal(test_framework, len(*cliServerPrivateNetwork.Ips), len(*sdkServerPrivateNetwork.Ips))

	for i := range *cliServerPrivateNetwork.Ips {
		assert.Equal(test_framework, (*cliServerPrivateNetwork.Ips)[i], (*sdkServerPrivateNetwork.Ips)[i])
	}

	assert.Equal(test_framework, cliServerPrivateNetwork.Dhcp, sdkServerPrivateNetwork.Dhcp)
	assert.Equal(test_framework, cliServerPrivateNetwork.StatusDescription, sdkServerPrivateNetwork.StatusDescription)
}
