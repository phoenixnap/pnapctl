package server

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	files "phoenixnap.com/pnap-cli/pnapctl/fileprocessor"
	"phoenixnap.com/pnap-cli/pnapctl/printer"
)

// Performs a Post request with a body containing a ServerCreate struct
// 		Receives a 200, 400, 500

// ServerCreate is the struct used as the body of the request
// to create a new server.
type ServerCreate struct {
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Public      bool     `json:"public" yaml:"public"`
	Os          string   `json:"os" yaml:"os"`
	TYPE        string   `json:"type" yaml:"type"`
	Location    string   `json:"location" yaml:"location"`
	SSHKeys     []string `json:"sshKeys" yaml:"sshKeys"`
}

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "create server"

var Full bool

// CreateServerCmd is the command for creating a server.
var CreateServerCmd = &cobra.Command{
	Use:          "server",
	Short:        "Create a new server.",
	Args:         cobra.ExactArgs(0),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Create a new server.

Requires a file (yaml or json) containing the information needed to create the server.`,
	Example: `# create a new server as described in server.yaml
pnapctl create server --filename ~/server.yaml

# server.yaml
name: "new-server"
description: "New server description"
public: true
os: "ubuntu/bionic"
type: "s1.c1.tiny"
location: "PHX"
sshKeys:
  - "ssh-rsa AAAAB3Nz...Fi9wrf+M7Q== test1@test"
  - "ssh-rsa AAAAB3Nz...dsWno-sa7nqt test2@test"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		files.ExpandPath(&Filename)

		data, err := files.ReadFile(Filename)

		if files.IsNotExist(err) {
			return ctlerrors.FileNotExistError(Filename)
		} else if err != nil {
			return ctlerrors.GenericNonRequestError(err.Error(), commandName)
		}

		// Marshal file into JSON using the struct
		var serverCreate ServerCreate

		structbyte, err := files.UnmarshalToJson(data, &serverCreate)

		if err != nil {
			return ctlerrors.GenericNonRequestError(err.Error(), commandName)
		}

		// Create the server
		response, err := client.MainClient.PerformPost("servers", bytes.NewBuffer(structbyte))

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName)
		}

		err = ctlerrors.GenerateErrorIfNot200(response, commandName)

		if err != nil {
			return err
		}

		return printer.PrintServerResponse(response.Body, false, Full, commandName)
	},
}

func init() {
	CreateServerCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	CreateServerCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateServerCmd.MarkFlagRequired("filename")
}
