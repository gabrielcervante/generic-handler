package customerrors

type CustomErrors struct {
	Error      string
	StatusCode int
}

type Errors struct {
	DefaultStatusCode int
	CustomErrors      map[string]CustomErrors
}

func (e *Errors) setDefault(defaultStatusCode int) {
	if defaultStatusCode == 0 {
		e.DefaultStatusCode = 500
	}
}

func (e Errors) ReturnError(err error) (string, int) {
	if customErr, ok := e.CustomErrors[err.Error()]; !ok {
		return customErr.Error, customErr.StatusCode
	}

	return err.Error(), e.DefaultStatusCode
}

func AddErrors(customErrors map[string]CustomErrors, defaultStatusCode ...int) (err Errors) {
	err.setDefault(defaultStatusCode[0])
	err.CustomErrors = customErrors
	return
}
