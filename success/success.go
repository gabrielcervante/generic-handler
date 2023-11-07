package success

type Success struct {
	DefaultStatusCode int
	SuccessStatusCode map[any]int
}

func (e *Success) setDefault(defaultStatusCode int) {
	if defaultStatusCode == 0 {
		e.DefaultStatusCode = 200
	}
}

func AddSuccesses(successStatusCode map[any]int, defaultStatusCode ...int) (sucess Success) {
	sucess.setDefault(defaultStatusCode[0])
	sucess.SuccessStatusCode = successStatusCode
	return
}

func (s Success) ReturnSuccess(message any) int {
	statusCode, ok := s.SuccessStatusCode[message]
	if !ok {
		return s.DefaultStatusCode
	}

	return statusCode
}
