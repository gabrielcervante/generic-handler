package handlers

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errorStatusCode    map[string]int
	errorReturnMessage map[string]string
)

type Errors struct {
	ErrorMessage       string
	ErrorStatusCode    int
	ErrorReturnMessage string
}

func (e Errors) validate() bool {
	if e.ErrorMessage == "" {
		return false
	}

	if e.ErrorStatusCode == 0 {
		return false
	}

	return true
}

func SaveErrors(customErrors ...Errors) error {
	errorStatusCode = make(map[string]int)
	errorReturnMessage = make(map[string]string)

	for i, err := range customErrors {

		if !err.validate() {
			errorStatusCode = nil
			errorReturnMessage = nil
			return errors.New(fmt.Sprintf("Sorry, your error in the position " + strconv.Itoa(i) + " have an error message or error status code empty."))
		}

		errorStatusCode[err.ErrorMessage] = err.ErrorStatusCode

		if err.ErrorReturnMessage != "" {
			errorReturnMessage[err.ErrorMessage] = err.ErrorReturnMessage
		}
	}

	return nil
}

func returnError(err string) (string, int) {
	statusCode, ok := errorStatusCode[err]
	if !ok {
		return err, 500
	}

	returnMessage, ok := errorReturnMessage[err]
	if !ok {
		return err, statusCode
	}

	return returnMessage, statusCode
}
