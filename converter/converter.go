package converter

import (
	"fmt"
)

var (
	convertWithError    map[string]any
	convertWithoutError map[string]any
)

func AddConvertWithError[I, O any](converters ...func(string) (O, error)) {
	var getType I

	for _, converter := range converters {
		if convertWithError == nil {
			convertWithError = make(map[string]any)
			convertWithError[fmt.Sprintf("%T", getType)] = converter
		}

		convertWithError[fmt.Sprintf("%T", getType)] = converter

	}
}

func AddconvertWithoutError[I, O any](converters ...func(string) O) {
	var getType I

	for _, converter := range converters {
		if convertWithoutError == nil {
			convertWithoutError = make(map[string]any)
			convertWithoutError[fmt.Sprintf("%T", getType)] = converter
		}

		convertWithoutError[fmt.Sprintf("%T", getType)] = converter
	}
}

func Convert[O any](input string) (any, error) {
	var getType O

	toConvertWithError, ok := convertWithError[fmt.Sprintf("%T", getType)]
	if ok {
		return toConvertWithError.(func(string) (O, error))(input)
	}

	toConvertWithoutError, ok := convertWithoutError[fmt.Sprintf("%T", getType)]
	if ok {
		return toConvertWithoutError.(func(string) error)(input), nil
	}

	return input, nil
}
