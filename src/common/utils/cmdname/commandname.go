package cmdname

import (
	"strings"

	"github.com/spf13/cobra"
)

var CommandName = "<NONE>"

func SetCommandName(cmd *cobra.Command) {
	CommandName = strings.Join(strings.Fields(cmd.CommandPath())[1:], " ")
}
