package commands

import (
	"fmt"
	"os"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/commands/create"
	"phoenixnap.com/pnap-cli/pnapctl/commands/delete"
	"phoenixnap.com/pnap-cli/pnapctl/commands/get"
	"phoenixnap.com/pnap-cli/pnapctl/commands/poweroff"
	"phoenixnap.com/pnap-cli/pnapctl/commands/poweron"
	"phoenixnap.com/pnap-cli/pnapctl/commands/reboot"
	"phoenixnap.com/pnap-cli/pnapctl/commands/reset"
	"phoenixnap.com/pnap-cli/pnapctl/commands/shutdown"
	"phoenixnap.com/pnap-cli/pnapctl/fileprocessor"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/configuration"
)

var (
	verbose bool
	cfgFile string

	rootCmd = &cobra.Command{
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
	if err := rootCmd.Execute(); err != nil {
		var _ = fmt.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	// add flags here when needed
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(reset.ResetCmd)
	rootCmd.AddCommand(delete.DeleteCmd)
	rootCmd.AddCommand(poweroff.PowerOffCmd)
	rootCmd.AddCommand(poweron.PowerOnCmd)
	rootCmd.AddCommand(shutdown.ShutdownCmd)
	rootCmd.AddCommand(reboot.RebootCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "change log level from Warn (default) to Debug.")

	cobra.OnInitialize(initConfig, setLoggingLevel)
}

func initConfig() {
	var configPath string
	envHome := os.Getenv("PNAPCTL_HOME")
	if envHome != "" && cfgFile == "" {
		cfgFile = envHome
	}

	if cfgFile != "" {
		// Use config file from the flag
		fileprocessor.ExpandPath(&cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		configPath = home + configuration.DefaultConfigPath
		// Search config in home directory with name "config" (without extension)
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Checks whether the config file exists, by attempting to cast the error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("A config file is required to run this program.\n" +
				"There are 3 approaches to specify the path of a configuration file (in order of priority)\n" +
				"\t1. --config flag: Specify the path and file name for the configuration file. (ex. pnapctl get servers --config=~/myconfig.yml\n" +
				"\t2. Environmental variable: Create an environmental variable called PNAPCTL_HOME specifying the path and filename.\n" +
				"\t3. Default: The default config file path is the home directory (" + configPath + "config.yaml)\n\n" +
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
		client.MainClient = client.NewHTTPClient(viper.GetString("clientId"), viper.GetString("clientSecret"))
	}
}

func setLoggingLevel() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
