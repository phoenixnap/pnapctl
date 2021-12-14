package servermodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerReserveToSDK(test_framework *testing.T) {
	cliModel := GenerateServerReserveCli()
	sdkModel := cliModel.toSDK()

	assert.Equal(test_framework, sdkModel.PricingModel, cliModel.PricingModel)
}
