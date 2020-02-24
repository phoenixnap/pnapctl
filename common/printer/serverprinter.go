package printer

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

type ShortServer struct {
	ID                 string   `header:"id"`
	Status             string   `header:"status"`
	Hostname           string   `header:"hostname"`
	Description        string   `header:"description"`
	PrivateIPAddresses []string `header:"Private Ips"`
	PublicIPAddresses  []string `header:"Public Ips"`
}

type LongServer struct {
	ID                 string   `header:"id"`
	Status             string   `header:"status"`
	Hostname           string   `header:"hostname"`
	Description        string   `header:"description"`
	PrivateIPAddresses []string `header:"Private Ips"`
	PublicIPAddresses  []string `header:"Public Ips"`
	Os                 string   `header:"os"`
	Type               string   `header:"type"`
	Location           string   `header:"location"`
	CPU                string   `header:"cpu"`
	RAM                string   `header:"ram"`
	Storage            string   `header:"storage"`
}

func PrintServerResponse(responseBody io.Reader, multiple bool, full bool, commandName string) error {
	body, err := ioutil.ReadAll(responseBody)

	if err != nil {
		return ctlerrors.CreateCLIError(ctlerrors.ResponseBodyReadFailure, commandName, err)
	}

	if full {
		if multiple {
			construct := &[]LongServer{}
			err = unmarshall(body, construct, commandName)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, len(*construct) == 0, commandName)
			}
		} else {
			construct := &LongServer{}
			err = unmarshall(body, construct, commandName)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, false, commandName)
			}
		}
	} else {
		if multiple {
			construct := &[]ShortServer{}
			err = unmarshall(body, construct, commandName)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, len(*construct) == 0, commandName)
			}
		} else {
			construct := &ShortServer{}
			err = unmarshall(body, construct, commandName)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, false, commandName)
			}
		}
	}

	// This err is the one outputted within the PrintOutput
	if err != nil {
		return err
	}

	return nil
}

// unmarshall will unmarshall a Json byte stream into the provided construct.
func unmarshall(body []byte, construct interface{}, commandName string) error {
	err := json.Unmarshal(body, &construct)
	if err != nil {
		return ctlerrors.CreateCLIError(ctlerrors.UnmarshallingErrorBody, commandName, err)
	}
	return nil
}
