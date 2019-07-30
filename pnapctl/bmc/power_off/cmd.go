package power_off

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
)

var P_OffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Powers off a specific server.",
	Long:  "Powers off a specific server.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-off':", len(args), "passed")
			panic("not-one-arg")
		}

		var resource = "servers/" + args[0] + "/actions/power-off"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			fmt.Println("Error while powering off server:", err)
			panic("power-off-error")
		} else if response.StatusCode == 409 {
			fmt.Println("Error: Conflict detected. Server is already powered-off.")
			panic("409-conflict")
		}
	},
}

func init() {
}
