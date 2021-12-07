package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Name        = "pnapctl"
	Version     = "dev"
	BuildDate   = "unknown"
	BuildCommit = "unknown"
)

var AppVersion AppVersionInfo

type AppVersionInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	BuildCommit  string `json:"revision"`
	GOVersion    string `json:"go_version"`
	BuiltAt      string `json:"built_at"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

func init() {
	AppVersion = AppVersionInfo{
		Name:         Name,
		Version:      Version,
		BuildCommit:  BuildCommit,
		GOVersion:    runtime.Version(),
		BuiltAt:      BuildDate,
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
	}
}

func (v *AppVersionInfo) String() string {
	version := fmt.Sprintf("Version:        %s\n", v.Version)
	version += fmt.Sprintf("Build commit:   %s\n", v.BuildCommit)
	version += fmt.Sprintf("Built date:     %s\n", v.BuiltAt)
	version += fmt.Sprintf("GO version:     %s\n", v.GOVersion)
	version += fmt.Sprintf("OS/Arch:        %s/%s\n", v.OS, v.Architecture)

	return version
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Display version and build information about pnapctl.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(AppVersion.String())
	},
}
