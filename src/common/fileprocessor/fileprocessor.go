package fileprocessor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

// MainFileProcessor is the main instance of FileProcessor that is used by the main code.
// Can be replaced in tests.
var MainFileProcessor FileProcessor = RealFileProcessor{}

// FileProcessor operates on files, including I/O.
type FileProcessor interface {
	ReadFile(filename string, commandName string) ([]byte, error)
}

// RealFileProcessor is an implementation of FileProcessor to be used by the main code.
type RealFileProcessor struct{}

// ReadFile Reads a file and processes any errors that may happen.
func (RealFileProcessor) ReadFile(filename string, commandName string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, ctlerrors.FileNotExistError(filename)
		} else {
			return nil, ctlerrors.CreateCLIError(ctlerrors.FileReading, commandName, err)
		}
	}

	return file, nil
}

// ReadFile is a shortcut function to using `MainFileProcessor` all the time.
func ReadFile(filename string, commandName string) ([]byte, error) {
	return MainFileProcessor.ReadFile(filename, commandName)
}

// unmarshal unmarshals a byte array into an object.
// The byte array can be either formatted as YAML or as JSON
func Unmarshal(data []byte, construct interface{}, commandName string) error {
	err := json.Unmarshal(data, construct)

	if err != nil {
		err = yaml.UnmarshalStrict(data, construct)

		if err != nil {
			return ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, commandName, err)
		}
	}
	return nil
}

// UnmarshalToJson unmarshals a byte array and marshals it back into JSON,
// using the struct that was passed as `construct`
func UnmarshalToJson(data []byte, construct interface{}, commandName string) ([]byte, error) {
	err := Unmarshal(data, construct, commandName)

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
