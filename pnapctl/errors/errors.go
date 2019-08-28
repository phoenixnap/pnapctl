package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
		Msg404: "404 NOT FOUND"
	}
}

func (r result) IfOk(message string) {
	r.Msg200 = message
}

func (r result) IfNotFound(message string) {
	r.Msg404 = message
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

	fmt.Println(bmcErr.Message, "\n", bmcErr.ValidationErrors)
	return errors.New(string(statusCode))
}
