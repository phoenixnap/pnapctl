package quotas

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestSubmitQuotaEditRequestSuccessYAML(test_framework *testing.T) {
	// setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, quotaEditLimitRequest)

	//prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(nil)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitQuotaEditRequestSuccessJSON(test_framework *testing.T) {
	//setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	ExpectFromFileSuccess(test_framework, json.Marshal, quotaEditLimitRequest)
	Filename = FILENAME

	//prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(nil)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitQuotaEditRequestFileProcessorFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	expectedErr := ExpectFromFileFailure(test_framework)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestSubmitQuotaEditRequestUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// prepare mocks
	ExpectFromFileUnmarshalFailure(test_framework)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestSubmitQuotaEditRequestYAMLUnmarshallingFailure(test_framework *testing.T) {
	// setup
	filecontents := []byte(`Limit: 45`)
	ExpectFromFileSuccess(test_framework, yaml.Marshal, filecontents)
	Filename = FILENAME

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, expectedErr, expectedErr.Error())
}

func TestSubmitQuotaEditClientFailure(test_framework *testing.T) {
	// setup
	editQuotaRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	ExpectFromFileSuccess(test_framework, yaml.Marshal, editQuotaRequest)
	Filename = FILENAME

	// prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(editQuotaRequest)).
		Return(testutil.TestError)

		// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
