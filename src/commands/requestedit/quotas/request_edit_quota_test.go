package quotas

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestSubmitQuotaEditRequestSuccessYAML(test_framework *testing.T) {
	// setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	yamlmarshal, _ := yaml.Marshal(quotaEditLimitRequest)

	Filename = FILENAME

	//prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitQuotaEditRequestSuccessJSON(test_framework *testing.T) {
	//setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	jsonmarshal, _ := json.Marshal(quotaEditLimitRequest)
	Filename = FILENAME

	//prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitQuotaEditRequestFileProcessorFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	ExpectFromFileFailure(test_framework)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := testutil.TestError

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestSubmitQuotaEditRequestUnmarshallingFailure(test_framework *testing.T) {
	// setup file with incorrect data
	filecontents := []byte(`limit 45`)
	Filename = FILENAME

	// prepare mocks
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestSubmitQuotaEditRequestYAMLUnmarshallingFailure(test_framework *testing.T) {
	// setup
	filecontents := []byte(`Limit: 45`)
	yamlmarshal, _ := yaml.Marshal(filecontents)
	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, expectedErr, expectedErr.Error())
}

func TestSubmitQuotaEditFileReadingFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestSubmitQuotaEditClientFailure(test_framework *testing.T) {
	// setup
	editQuotaRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	yamlmarshal, _ := yaml.Marshal(editQuotaRequest)
	Filename = FILENAME

	// prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(editQuotaRequest)).
		Return(testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
