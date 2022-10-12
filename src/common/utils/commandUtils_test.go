package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestIs200Success(test_framework *testing.T) {
	assert.Equal(test_framework, is2xxSuccessful(200), true)
}

func TestIs201Success(test_framework *testing.T) {
	assert.Equal(test_framework, is2xxSuccessful(201), true)
}

func TestIs202Success(test_framework *testing.T) {
	assert.Equal(test_framework, is2xxSuccessful(202), true)
}

func TestIs500Fail(test_framework *testing.T) {
	assert.Equal(test_framework, is2xxSuccessful(500), false)
}

func TestCheckForErrors_nilResponse(test_framework *testing.T) {
	assert.Equal(test_framework, CheckErrs(nil, nil), nil)
}

func TestCheckForErrors_400Response(test_framework *testing.T) {
	response := WithResponse(400, WithBody(nil))

	expectedErr := ctlerrors.HandleBMCError(response)
	err := CheckErrs(response, nil)

	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCheckForErrors_SDKErr(test_framework *testing.T) {
	err := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)
	generatedError := CheckErrs(nil, err)

	assert.EqualError(test_framework, err, generatedError.Error())
}
