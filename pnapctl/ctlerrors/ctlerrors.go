package ctlerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type result struct {
	Msg200 string
	Msg404 string
}

type BMCError struct {
	Message          string
	ValidationErrors []string
}

func Result() result {
	return result{
		Msg200: "200 OK",
		Msg404: "404 NOT FOUND",
	}
}

func (r result) IfOk(message string) result {
	r.Msg200 = message
	return r
}

func (r result) IfNotFound(message string) result {
	r.Msg404 = message
	return r
}

func (b BMCError) String() string {
	if len(b.ValidationErrors) == 0 {
		return b.Message
	} else {
		return b.Message + "\n" + strings.Join(b.ValidationErrors, "\n")
	}
}

func TryUnmarshal(body []byte, construct *interface{}) error {
	err := json.Unmarshal(body, &construct)

	if err == nil {
		return nil
	}

	err_b := json.Unmarshal(body, &BMCError{})
	if err_b != nil {
		return errors.New("body-is-error")
	} else {
		return err
	}
}

func (r result) UseResponse(response *http.Response) error {
	statusCode := response.StatusCode

	if statusCode == 200 {
		fmt.Println(r.Msg200)
		return nil
	} else if statusCode == 404 {
		fmt.Println(r.Msg404)
		return errors.New("404")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return errors.New("no-body-and-not-400-or-200")
	}

	bmcErr := BMCError{}
	json.Unmarshal(body, &bmcErr)

	fmt.Println(bmcErr)
	return errors.New(string(statusCode))
}
