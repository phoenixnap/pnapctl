package pnapctl

import (
	"fmt"
	"os"

	"phoenixnap.com/pnap-cli/pnapctl/client"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"phoenixnap.com/pnap-cli/pnapctl/bmc"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "pnapctl",
	Short: "Short Desc",
	Long:  "Longer Desc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

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
	rootCmd.AddCommand(bmc.BmcCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default \"$HOME/.pnap.yaml\")")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "pnap" (withou extension)
		viper.AddConfigPath(home)
		viper.SetConfigName("pnap")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	} else if viper.GetString("clientId") == "" || viper.GetString("clientSecret") == "" {
		fmt.Println("Client ID and Client Secret in config file should not be empty")
		os.Exit(1)
	} else {
		client.MainClient = client.NewHTTPClient(viper.GetString("clientId"), viper.GetString("clientSecret"))
	}
}
