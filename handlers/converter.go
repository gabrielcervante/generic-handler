package handlers

import (
	"fmt"
	"strconv"
)

func convert[I any](input string) (any, error) {
	var getType I
	switch fmt.Sprintf("%T", getType) {
	case "int":
		return strconv.Atoi(input)
	case "bool":
		return strconv.ParseBool(input)
	default:
		return input, nil
	}
}
