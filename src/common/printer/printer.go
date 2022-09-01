package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	iter "phoenixnap.com/pnapctl/common/utils/iterutils"
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

// Used in order to pass a 'prepare' function that uses 'full' into [iterutils.Map].
//
//	func PrepareServer(srv Server, full bool) interface{} {
//		// ...
//	}
//
//	serversToPrint := iterutils.Map(sdkServers, WithFull(true, PrepareServer))
func withFull[T any](full bool, mapper iter.Currier[T, bool, interface{}]) iter.Mapper[T, interface{}] {
	return iter.Curry(mapper, full)
}

// Prepares a OneOf response for printing. It works by:
//  1. Mapping each item by using the mapper passed.
//  2. Filtering out all 'nil' values (aka unrecognized types)
//  3. Dereferencing every value (since we removed all 'nil' values)
//
// Step 3. is required as a pointer to a table struct will not print correctly.
func prepareOneOfWith[In any](in []In, mapper iter.Mapper[In, interface{}]) (out []interface{}) {
	out = iter.Map(in, mapper)
	out = iter.Filter(out, iter.Not(iter.IsNil))
	return iter.Deref(out)
}
