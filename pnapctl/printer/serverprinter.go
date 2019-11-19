package printer

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

type ShortServer struct {
	ID                 string   `header:"id"`
	Status             string   `header:"status"`
	Name               string   `header:"name"`
	Description        string   `header:"description"`
	PrivateIPAddresses []string `header:"Private Ips"`
	PublicIPAddresses  []string `header:"Public Ips"`
}

type LongServer struct {
	ID                 string   `header:"id"`
	Status             string   `header:"status"`
	Name               string   `header:"name"`
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
		return ctlerrors.GenericNonRequestError(ctlerrors.ResponseBodyReadFailure, commandName)
	}

	if full {
		if multiple {
			construct := &[]LongServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, len(*construct) == 0)
			}
		} else {
			construct := &LongServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, false)
			}
		}
	} else {
		if multiple {
			construct := &[]ShortServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, len(*construct) == 0)
			}
		} else {
			construct := &ShortServer{}
			err = unmarshall(body, construct)
			if err == nil {
				err = MainPrinter.PrintOutput(construct, false)
			}
		}
	}

	// This err is the one outputted within the PrintOutput
	if err != nil {
		return ctlerrors.GenericNonRequestError(err.Error(), commandName)
	}

	return nil
}

// unmarshall will unmarshall a Json byte stream into the provided construct.
func unmarshall(body []byte, construct interface{}) error {
	err := json.Unmarshal(body, &construct)
	if err != nil {
		return errors.New(ctlerrors.UnmarshallingErrorBody)
	}
	return nil
}
