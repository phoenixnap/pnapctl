package testutil

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertNilEquality(test_framework *testing.T, varName string, cliVar interface{}, sdkVar interface{}) bool {
	if cliVar == nil || reflect.ValueOf(cliVar).IsNil() {
		assert.Nil(test_framework, sdkVar, "(value: "+varName+") CLI's value is nil, but not SDK's value.")
		return false
	} else if sdkVar == nil || reflect.ValueOf(sdkVar).IsNil() {
		assert.Nil(test_framework, cliVar, "(value: "+varName+") SDK's value is nil, but not CLI's value.")
		return false
	}

	return true
}

func AssertIsType[T any](test_framework *testing.T, item interface{}) {
	var t T
	assert.Equal(
		test_framework,
		fmt.Sprintf("%T", item),
		fmt.Sprintf("%T", t),
	)
}
