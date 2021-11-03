package quotas

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "edit quota"

var Full bool

// CreateServerCmd is the command for creating a server.
var EditQuotaCmd = &cobra.Command{
	Use:   "quota [QUOTA_ID]",
	Short: "Subit a quota modification request.",
	Args:  cobra.ExactArgs(1),
	//Aliases:      []string{"quota"},
	SilenceUsage: true,
	Long: `Submit a quota modification request.

Requires a file (yaml or json) containing the information needed to submit the request.`,
	Example: `# modify an existing quota as per quotaModificationRequest.yaml
pnapctl edit quota  --filename ~/quotaModificationRequest.yaml

# quotaModificationRequest.yaml
limit: 75
reason: "My current limit is not enough."`,
	RunE: func(cmd *cobra.Command, args []string) error {
		quotaEditRequest, err := models.CreateQuotaEditRequestFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		httpResponse, err := bmcapi.Client.QuotaEditById(args[0], *quotaEditRequest)

		if err != nil {
			// TODO - Validate way of processing errors.
			return err
		} else if httpResponse.StatusCode != 202 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println("Quota Edit Limit Request Accepted.")
		return nil
	},
}

func init() {
	EditQuotaCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	EditQuotaCmd.MarkFlagRequired("filename")
}
