package converter

type Converter[O any] struct {
	ConvertWithError    func(string) (O, error)
	ConvertWithoutError func(string) O
}

func (c Converter[O]) Convert(input string) (any, error) {
	if c.ConvertWithError != nil {
		return c.ConvertWithError(input)
	}

	if c.ConvertWithoutError != nil {
		return c.ConvertWithoutError(input), nil
	}

	return input, nil
}
