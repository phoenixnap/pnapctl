package printer

import (
	"bytes"
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
	PrintOutput(body []byte, construct interface{}) error
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
// The first output parameter is only used by the table printer to show how
// many rows were printed. The second output parameter specifies any errors.
func (m BodyPrinter) PrintOutput(body []byte, construct interface{}) error {
	err := json.Unmarshal(body, &construct)

	if err != nil {
		return errors.New(ctlerrors.UnmarshallingInPrinter)
	}

	if OutputFormat == "json" {
		return printJSON(body)
	} else if OutputFormat == "yaml" {
		return printYAML(construct)
	} else {
		// default to table
		rows := printTable(construct, m.Tableprinter)
		if rows == -1 {
			return errors.New(ctlerrors.TablePrinterFailure)
		}
		return nil
	}
}

// Attempts to print in JSON via formatting a byte array.
func printJSON(body []byte) error {
	var dat bytes.Buffer
	json.Indent(&dat, body, "", "    ")
	fmt.Println(string(dat.Bytes()))

	return nil
}

// Attempts to print in YAML via marshalling.
func printYAML(body interface{}) error {
	b, err := yaml.Marshal(body)

	if err != nil {
		return errors.New(ctlerrors.MarshallingInPrinter)
	}

	fmt.Println(string(b))

	return nil
}

// Attempts to print the struct as a table.
func printTable(body interface{}, tblprinter *tableprinter.Printer) int {
	return tblprinter.Print(body)
}
