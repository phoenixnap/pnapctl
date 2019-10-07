package printer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

// The main printer used by the application.
// Can be swapped by a Mock for testing.
var MainPrinter = NewBodyPrinter()

// The format to print the object in.
var OutputFormat string

type Printer interface {
	PrintOutput(construct interface{}, isEmpty bool) error
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
func (m BodyPrinter) PrintOutput(construct interface{}, isEmpty bool) error {
	if OutputFormat == "json" {
		return printJSON(construct)
	} else if OutputFormat == "yaml" {
		return printYAML(construct)
	} else {
		if isEmpty {
			// We cannot print a table if we don't have at least the headers.
			fmt.Println("No data found")
			return nil
		} else {
			// default to table
			return printTable(construct, m.Tableprinter)
		}
	}
}

// printJSON attempts to print in JSON via marshalling.
func printJSON(body interface{}) error {
	b, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		return errors.New(ctlerrors.MarshallingInPrinter)
	}

	fmt.Println(string(b))
	return nil
}

// printYAML attempts to print in YAML via marshalling.
func printYAML(body interface{}) error {
	b, err := yaml.Marshal(body)
	if err != nil {
		return errors.New(ctlerrors.MarshallingInPrinter)
	}

	fmt.Println(string(b))
	return nil
}

// Attempts to print the struct as a table.
func printTable(body interface{}, tblprinter *tableprinter.Printer) error {
	rows := tblprinter.Print(body)

	if rows == -1 {
		return errors.New(ctlerrors.TablePrinterFailure)
	}

	return nil
}
