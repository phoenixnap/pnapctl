package quotas

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "request-edit quota"

// RequestEditQuotaCmd is the command for requesting a quota modification.
var RequestEditQuotaCmd = &cobra.Command{
	Use:          "quota QUOTA_ID",
	Short:        "Submit a quota modification request.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Submit a quota modification request.

Requires a file (yaml or json) containing the information needed to submit a quota edit request.`,
	Example: `# Submit an edit request on an existing quota as per requestEditQuota.yaml
pnapctl request-edit quota <QUOTA_ID> --filename <FILE_PATH>

# requestEditQuota.yaml
limit: 75
reason: My current limit is not enough.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		quotaEditRequest, err := models.CreateRequestFromFile[bmcapisdk.QuotaEditLimitRequest](Filename, commandName)
		if err != nil {
			return err
		}

		httpResponse, err := bmcapi.Client.QuotaEditById(args[0], *quotaEditRequest)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			fmt.Println("Quota Edit Limit Request Accepted.")
			return nil
		}
	},
}

func init() {
	RequestEditQuotaCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	RequestEditQuotaCmd.MarkFlagRequired("filename")
}
