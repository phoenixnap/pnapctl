package fileprocessor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

// MainFileProcessor is the main instance of FileProcessor that is used by the main code.
// Can be replaced in tests.
var MainFileProcessor FileProcessor = RealFileProcessor{}

// FileProcessor operates on files, including I/O.
type FileProcessor interface {
	ReadFile(filename string) ([]byte, error)
}

// RealFileProcessor is an implementation of FileProcessor to be used by the main code.
type RealFileProcessor struct{}

// ReadFile Reads a file and processes any errors that may happen.
func (RealFileProcessor) ReadFile(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(ctlerrors.FileDoesNotExist)
		} else {
			return nil, errors.New(ctlerrors.FileReading)
		}
	}

	return file, nil
}

// IsNotExist checks whether an error is FileDoesNotExist
func IsNotExist(err error) bool {
	return err != nil && err.Error() == ctlerrors.FileDoesNotExist
}

// ReadFile is a shortcut function to using `MainFileProcessor` all the time.
func ReadFile(filename string) ([]byte, error) {
	return MainFileProcessor.ReadFile(filename)
}

// unmarshal unmarshals a byte array into an object.
// The byte array can be either formatted as YAML or as JSON
func unmarshal(data []byte, construct interface{}) error {
	err := json.Unmarshal(data, construct)

	if err != nil {
		err = yaml.Unmarshal(data, construct)

		if err != nil {
			return errors.New(ctlerrors.UnmarshallingInFileProcessor)
		}
	}
	return nil
}

// UnmarshalToJson unmarshals a byte array and marshals it back into JSON,
// using the struct that was passed as `construct`
func UnmarshalToJson(data []byte, construct interface{}) ([]byte, error) {
	err := unmarshal(data, construct)

	if err != nil {
		return nil, err
	}

	s, _ := json.Marshal(construct)

	return s, nil
}

// ExpandPath expands the path sent using the shell. Cross-compatible.
func ExpandPath(path *string) {
	// Uses echo to let shell expand the path
	var cmd *exec.Cmd
	var endlineChar string

	if runtime.GOOS == "windows" {
		endlineChar = "\r\n"
		cmd = exec.Command("cmd", "/c", fmt.Sprintf("echo %s", *path))
	} else {
		endlineChar = "\n"
		cmd = exec.Command("sh", "-c", fmt.Sprintf("''echo %s''", *path))
	}

	// Captures output of echo
	byteoutput, _ := cmd.Output()

	// Sets path to output - removing trailing newline.
	*path = strings.TrimSuffix(string(byteoutput), endlineChar)
}
