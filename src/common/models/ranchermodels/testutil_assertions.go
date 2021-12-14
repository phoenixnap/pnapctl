package ranchermodels

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNilEquality(test_framework *testing.T, varName string, cliVar interface{}, sdkVar interface{}) bool {
	if cliVar == nil || reflect.ValueOf(cliVar).IsNil() {
		assert.Nil(test_framework, sdkVar, "(value: "+varName+") CLI's value is nil, but not SDK's value.")
		return true
	} else if sdkVar == nil || reflect.ValueOf(sdkVar).IsNil() {
		assert.Nil(test_framework, cliVar, "(value: "+varName+") SDK's value is nil, but not CLI's value.")
		return true
	}

	return false
}
