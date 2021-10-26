package printer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/common/ctlerrors"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"

	"phoenixnap.com/pnap-cli/common/models"
)

// The main printer used by the application.
// Can be swapped by a Mock for testing.
var MainPrinter = NewBodyPrinter()

// The format to print the object in.
var OutputFormat string

type Printer interface {
	PrintOutput(construct interface{}, isEmpty bool, commandName string) error
}

type BodyPrinter struct {
	Tableprinter *tableprinter.Printer
}

func NewBodyPrinter() Printer {
	return BodyPrinter{
		Tableprinter: tableprinter.New(os.Stdout),
	}
}

// PrintOutput prints the construct passed according to the format.
// The output parameter specifies whether any errors were encountered during printing
func (m BodyPrinter) PrintOutput(construct interface{}, isEmpty bool, commandName string) error {
	if OutputFormat == "json" {
		return printJSON(construct, commandName)
	} else if OutputFormat == "yaml" {
		return printYAML(construct, commandName)
	} else {
		if isEmpty {
			// We cannot print a table if we don't have at least the headers.
			fmt.Println("No data found")
			return nil
		} else {
			// default to table
			return printTable(construct, m.Tableprinter, commandName)
		}
	}
}

// printJSON attempts to print in JSON via marshalling.
func printJSON(body interface{}, commandName string) error {
	b, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		return ctlerrors.CreateCLIError(ctlerrors.MarshallingInPrinter, commandName, err)
	}

	fmt.Println(string(b))
	return nil
}

// printYAML attempts to print in YAML via marshalling.
func printYAML(body interface{}, commandName string) error {
	b, err := yaml.Marshal(body)
	if err != nil {
		return ctlerrors.CreateCLIError(ctlerrors.MarshallingInPrinter, commandName, err)
	}

	fmt.Println(string(b))
	return nil
}

// Attempts to print the struct as a table.
func printTable(body interface{}, tblprinter *tableprinter.Printer, commandName string) error {
	tblprinter.RowCharLimit = 23
	rows := tblprinter.Print(body)

	emptyBody := false
	switch x := body.(type) {
	case []interface{}:
		emptyBody = len(x) == 0
	default:
		emptyBody = x == nil
	}

	if emptyBody {
		fmt.Println("Nothing found.")
		return nil
	} else if rows == -1 {
		return ctlerrors.CreateCLIError(ctlerrors.MarshallingInPrinter, commandName, nil)
	}

	return nil
}

func PrintServerResponse(server bmcapi.Server, full bool, commandName string) error {
	if full {
		return MainPrinter.PrintOutput(models.ToFullServer(server), false, commandName)
	} else {
		return MainPrinter.PrintOutput(models.ToShortServer(server), false, commandName)
	}
}

func PrintServerListResponse(servers []bmcapi.Server, full bool, commandName string) error {
	var serverList []interface{}

	if full {
		for _, bmcServer := range servers {
			serverList = append(serverList, models.ToFullServer(bmcServer))
		}
	} else {
		for _, bmcServer := range servers {
			serverList = append(serverList, models.ToShortServer(bmcServer))
		}
	}

	return MainPrinter.PrintOutput(serverList, false, commandName)
}
