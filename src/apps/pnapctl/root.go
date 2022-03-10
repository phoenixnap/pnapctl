package commands

import (
	"fmt"
	"os"
	"phoenixnap.com/pnapctl/commands/deprovision"
	"phoenixnap.com/pnapctl/common/client/ip"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/create"
	"phoenixnap.com/pnapctl/commands/delete"
	"phoenixnap.com/pnapctl/commands/get"
	"phoenixnap.com/pnapctl/commands/patch"
	"phoenixnap.com/pnapctl/commands/poweroff"
	"phoenixnap.com/pnapctl/commands/poweron"
	"phoenixnap.com/pnapctl/commands/reboot"
	"phoenixnap.com/pnapctl/commands/requestedit"
	"phoenixnap.com/pnapctl/commands/reserve"
	"phoenixnap.com/pnapctl/commands/reset"
	"phoenixnap.com/pnapctl/commands/shutdown"
	"phoenixnap.com/pnapctl/commands/tag"
	"phoenixnap.com/pnapctl/commands/update"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client/audit"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/client/tags"
	"phoenixnap.com/pnapctl/common/fileprocessor"
	configuration "phoenixnap.com/pnapctl/configs"
)

const HOME_ENV_VAR = "PNAPCTL_HOME"
const DEFAULT_CFG_NAME = "config"

var (
	verbose bool
	cfgFile string

	RootCmd = &cobra.Command{
		Use:   "pnapctl",
		Short: "pnapctl creates new and manages existing bare metal servers.",
		Long: `pnapctl creates new and manages existing bare metal servers provided by the phoenixNAP Bare Metal Cloud service.
	
	Find More information at: ` + configuration.KnowledgeBaseURL,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}
)

// Execute adds all child commands to the root command, setting flags appropriately.
// Called by main.main(), only needing to happen once.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		var _ = fmt.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	// add flags here when needed
	RootCmd.AddCommand(get.GetCmd)
	RootCmd.AddCommand(create.CreateCmd)
	RootCmd.AddCommand(update.UpdateCmd)
	RootCmd.AddCommand(patch.PatchCmd)
	RootCmd.AddCommand(reset.ResetCmd)
	RootCmd.AddCommand(delete.DeleteCmd)
	RootCmd.AddCommand(poweroff.PowerOffCmd)
	RootCmd.AddCommand(poweron.PowerOnCmd)
	RootCmd.AddCommand(shutdown.ShutdownCmd)
	RootCmd.AddCommand(reboot.RebootCmd)
	RootCmd.AddCommand(reserve.ReserveCmd)
	RootCmd.AddCommand(deprovision.DeprovisionCmd)
	RootCmd.AddCommand(version.VersionCmd)
	RootCmd.AddCommand(requestedit.RequestEditCmd)
	RootCmd.AddCommand(tag.TagCmd)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "change log level from Warn (default) to Debug.")

	cobra.OnInitialize(initConfig, setLoggingLevel)
}

func initConfig() {
	var defaultHomeDir string

	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defaultHomeDir = home + configuration.DefaultConfigPath

	if cfgFile != "" {
		// Use config file from the flag
		fileprocessor.ExpandPath(&cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Use the configured Home from env var
		cfgPath := os.Getenv(HOME_ENV_VAR)

		if cfgPath == "" {
			// Use the default home
			cfgPath = defaultHomeDir
		}

		// Search config in home directory (without extension)
		viper.AddConfigPath(cfgPath)
		viper.SetConfigName(DEFAULT_CFG_NAME)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Checks whether the config file exists, by attempting to cast the error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("A config file is required to run this program.\n" +
				"There are 3 approaches to specify the path of a configuration file (in order of priority)\n" +
				"\t1. --config flag: Specify the path and file name for the configuration file (ex. pnapctl get servers --config=~/myconfig.yaml)\n" +
				"\t2. Environmental variable: Create an environmental variable called " + HOME_ENV_VAR + " specifying the path containing the configuration file (" + DEFAULT_CFG_NAME + ".yaml)\n" +
				"\t3. Default: The default config file path is the home directory (" + defaultHomeDir + "config.yaml)\n\n" +
				"Find More information at: " + configuration.KnowledgeBaseURL + "\n\n" +
				"The following shows a sample config file:\n\n" +
				"# =====================================================\n" +
				"# Sample yaml config file\n" +
				"# =====================================================\n\n" +
				"# Authentication\n" +
				"clientId: <enter your client id>\n" +
				"clientSecret: <enter your client secret>")
		} else {
			fmt.Println("Error reading config file:", err)
		}

		os.Exit(1)
	} else if viper.GetString("clientId") == "" || viper.GetString("clientSecret") == "" {
		fmt.Println("Client ID and Client Secret in config file should not be empty")
		os.Exit(1)
	} else {
		clientId := viper.GetString("clientId")
		clientSecret := viper.GetString("clientSecret")

		customBmcApiHostname := viper.GetString("bmcApiHostname")
		customRancherHostname := viper.GetString("rancherHostname")
		customAuditHostname := viper.GetString("auditHostname")
		customTagsHostname := viper.GetString("tagsHostname")
		customNetworksHostname := viper.GetString("networksHostname")
		customIpHostname := viper.GetString("ipHostname")
		customTokenUrl := viper.GetString("tokenURL")

		bmcapi.Client = bmcapi.NewMainClient(clientId, clientSecret, customBmcApiHostname, customTokenUrl)
		rancher.Client = rancher.NewMainClient(clientId, clientSecret, customRancherHostname, customTokenUrl)
		audit.Client = audit.NewMainClient(clientId, clientSecret, customAuditHostname, customTokenUrl)
		tags.Client = tags.NewMainClient(clientId, clientSecret, customTagsHostname, customTokenUrl)
		networks.Client = networks.NewMainClient(clientId, clientSecret, customNetworksHostname, customTokenUrl)
		ip.Client = ip.NewMainClient(clientId, clientSecret, customIpHostname, customTokenUrl)
	}
}

func setLoggingLevel() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
