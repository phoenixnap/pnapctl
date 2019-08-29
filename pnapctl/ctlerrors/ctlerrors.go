package ctlerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (r result) UseResponse(response *http.Response) error {
	statusCode := response.StatusCode

	if statusCode == 200 {
		fmt.Println(r.Msg200)
		return nil
	} else if statusCode == 404 {
		fmt.Println(r.Msg404)
		return errors.New("404")
	}

	if response.Body == nil {
		return errors.New("no-body-and-not-400-or-200")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	bmcErr := BMCError{}
	err = json.Unmarshal(body, &bmcErr)

	if err != nil {
		return err
	}

	fmt.Println("Error:", bmcErr)
	return errors.New(strconv.Itoa(statusCode))
}
