package example

import (
	"context"
	"errors"

	"github.com/gabrielcervante/gin-handler/handlers"
	"github.com/gin-gonic/gin"
)

func Example() {
	err := handlers.SaveErrors(handlers.Errors{ErrorMessage: "Sorry error in string", ErrorStatusCode: 500, ErrorReturnMessage: "Something went wrong"},
		handlers.Errors{ErrorMessage: "Sorry error in int", ErrorStatusCode: 404})
	if err != nil {
		return
	}

	err = handlers.SaveSuccesses(handlers.Successes{SuccessMessage: "", SuccessStatusCode: 201},
		handlers.Successes{SuccessMessage: customType{}, SuccessStatusCode: 200})
	if err != nil {
		return
	}

	r := gin.Default()

	r.GET("/string/:str", handlers.Newhandler[string, string](nil, nil).HandleParam("str", respondMessageString))
	r.GET("/str", handlers.Newhandler[string, string](nil, nil).HandleParam("string", errorMessageString))
	r.GET("/int", handlers.Newhandler[int, int](nil, nil).HandleQuery("integer", errorMessageInt))

	postHandlers := handlers.Newhandler[customType, customType](nil, nil)

	r.POST("/custom", postHandlers.HandleJSON(respondMessageCustomType))

	err = r.Run()
	if err != nil {
		return
	}
}

func errorMessageInt(ctx context.Context, i int) (int, error) {
	return i, errors.New("Sorry error in int")
}

func errorMessageString(ctx context.Context, s string) (string, error) {
	return s, errors.New("Sorry error in string")
}

func respondMessageString(ctx context.Context, s string) (string, error) {
	return s, nil
}

type customType struct {
	Str string `json:"str"`
}

func respondMessageCustomType(ctx context.Context, s customType) (customType, error) {
	return s, nil
}
