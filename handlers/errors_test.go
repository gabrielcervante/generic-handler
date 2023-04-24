package handlers

import (
	"fmt"
	"testing"
)

func TestSaveErrors(t *testing.T) {
	fmt.Println(SaveErrors(Errors{ErrorMessage: "Slv", ErrorStatusCode: 404, ErrorReturnMessage: "ss"}, Errors{ErrorStatusCode: 400, ErrorReturnMessage: "affs"}))
	fmt.Println(returnError("Slv"))
}
