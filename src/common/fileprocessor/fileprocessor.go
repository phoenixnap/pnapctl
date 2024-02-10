package fileprocessor

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"sigs.k8s.io/yaml"
	"github.com/rs/zerolog/log"
)

// MainFileProcessor is the main instance of FileProcessor that is used by the main code.
// Can be replaced in tests.
var MainFileProcessor FileProcessor = RealFileProcessor{}

// FileProcessor operates on files, including I/O.
type FileProcessor interface {
	ReadFile(filename string) ([]byte, error)
	SaveFile(filename string, file *os.File) (error)
}

// RealFileProcessor is an implementation of FileProcessor to be used by the main code.
type RealFileProcessor struct{}

// ReadFile Reads a file and processes any errors that may happen.
func (RealFileProcessor) ReadFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, ctlerrors.FileNotExistError(filename)
		} else {
			return nil, ctlerrors.CreateCLIError(ctlerrors.FileReading, err)
		}
	}

	return file, nil
}

func (RealFileProcessor) SaveFile(filename string, file *os.File) error {

	// Get the file information number of bytes
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Prepare the byte array to store the content
	data := make([]byte, fileInfo.Size())

	// Read the content of the file to byte array
	_, err = file.Read(data)
	if err != nil {
		return err
	}

	// Open the file with write-only permission and create it if it doesn't exist
	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	// Close the file when function is completed
	defer file.Close()

	// Write the content to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	log.Info().Msgf("File '%s' downloaded successfully.\n", filename)
	return nil
}

// ReadFile is a shortcut function to using `MainFileProcessor` all the time.
func ReadFile(filename string) ([]byte, error) {
	return MainFileProcessor.ReadFile(filename)
}

// SaveFile is a shortcut function to using `MainFileProcessor` all the time.
func SaveFile(filename string, file *os.File) (error) {
	return MainFileProcessor.SaveFile(filename, file)
}

// unmarshal unmarshals a byte array into an object.
// The byte array can be either formatted as YAML or as JSON
func Unmarshal(data []byte, construct interface{}) error {
	err := json.Unmarshal(data, construct)

	if err != nil {
		err = yaml.UnmarshalStrict(data, construct)

		if err != nil {
			return ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)
		}
	}
	return nil
}

// UnmarshalToJson unmarshals a byte array and marshals it back into JSON,
// using the struct that was passed as `construct`
func UnmarshalToJson(data []byte, construct interface{}) ([]byte, error) {
	err := Unmarshal(data, construct)

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
