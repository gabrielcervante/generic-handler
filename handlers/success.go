package handlers

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	successStatusCode map[string]int
)

type Successes struct {
	SuccessMessage    any
	SuccessStatusCode int
}

func (e Successes) validate() bool {
	if e.SuccessMessage == nil {
		return false
	}

	if e.SuccessStatusCode == 0 {
		return false
	}

	return true
}

func SaveSuccesses(success ...Successes) error {
	successStatusCode = make(map[string]int)

	for i, s := range success {
		if !s.validate() {
			successStatusCode = nil
			return errors.New(fmt.Sprintf("Sorry, your success in the position " + strconv.Itoa(i) + " have an success message or success status code empty."))
		}

		successStatusCode[fmt.Sprintf("%T", s.SuccessMessage)] = s.SuccessStatusCode
	}

	return nil
}

func returnSuccess(message any) (any, int) {
	statusCode, ok := successStatusCode[fmt.Sprintf("%T", message)]
	if !ok {
		return message, 200
	}

	return message, statusCode
}
