package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

// The main printer used by the application.
// Can be swapped by a Mock for testing.
var MainPrinter = NewBodyPrinter()

// The format to print the object in.
var OutputFormat string

type Printer interface {
	PrintOutput(construct interface{}, commandName string) error
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
func (m BodyPrinter) PrintOutput(construct interface{}, commandName string) error {
	if OutputFormat == "json" {
		return printJSON(construct, commandName)
	} else if OutputFormat == "yaml" {
		return printYAML(construct, commandName)
	} else {
		return printTable(construct, m.Tableprinter, commandName)
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

	switch reflect.TypeOf(body).Kind() {
	case reflect.Slice:
		list := reflect.ValueOf(body)
		emptyBody = list.Len() == 0
	default:
		emptyBody = body == nil
	}

	if emptyBody {
		fmt.Println("No data found.")
		return nil
	} else if rows == -1 {
		return ctlerrors.CreateCLIError(ctlerrors.MarshallingInPrinter, commandName, nil)
	}

	return nil
}

func OutputIsTable() bool {
	return OutputFormat != "json" && OutputFormat != "yaml"
}
