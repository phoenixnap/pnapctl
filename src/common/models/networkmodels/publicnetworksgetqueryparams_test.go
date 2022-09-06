package networkmodels

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestPublicNetworksGetQueryParamsValid_Success(test_framework *testing.T) {
	iterutils.Each(allowedLocations, func(location string) {
		queryParams, err := NewPublicNetworksGetQueryParams(location)

		assert.NoError(test_framework, err)
		assert.Equal(test_framework, &location, queryParams.Location)
	})
}

func TestPublicNetworksGetQueryParamsEmpty_Success(test_framework *testing.T) {
	queryParams, err := NewPublicNetworksGetQueryParams("")

	assert.NoError(test_framework, err)
	assert.Nil(test_framework, queryParams.Location)
}

func TestPublicNetworksGetQueryParamsInvalid_Error(test_framework *testing.T) {
	queryParams, err := NewPublicNetworksGetQueryParams("NOTVALID")

	assert.Nil(test_framework, queryParams)
	assert.EqualError(test_framework, err, fmt.Sprintf("location '%s' is invalid. Allowed values are %v", "NOTVALID", allowedLocations))
}
