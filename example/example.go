package example

import (
	"context"
	"errors"

	customerrors "github.com/gabrielcervante/handler/custom_errors"
	"github.com/gabrielcervante/handler/handlers"
	"github.com/gabrielcervante/handler/success"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func FiberExample() {
	customErrors := customerrors.AddErrors(map[string]customerrors.CustomErrors{
		"Sorry error in string": {Error: "Something went wrong", StatusCode: 500},
		"Sorry error in int":    {StatusCode: 404},
	}, 505)

	successes := success.AddSuccesses(map[any]int{
		"":           201,
		customType{}: 200,
	}, 505)

	r := fiber.New()

	r.Get("/string/:str", handlers.NewFiberHandler[string, string](customErrors, successes).HandleParam("str", respondMessageString))
	r.Get("/str", handlers.NewFiberHandler[string, string](customErrors, successes).HandleParam("string", errorMessageString))
	r.Get("/int", handlers.NewFiberHandler[int, int](customErrors, successes).HandleQuery("integer", errorMessageInt))

	postHandlers := handlers.NewFiberHandler[customType, customType](customErrors, successes)

	r.Post("/custom", postHandlers.HandleJSON(respondMessageCustomType))

	err := r.Listen(":8080")
	if err != nil {
		return
	}
}

func GinExample() {
	customErrors := customerrors.AddErrors(map[string]customerrors.CustomErrors{
		"Sorry error in string": {Error: "Something went wrong", StatusCode: 500},
		"Sorry error in int":    {StatusCode: 404},
	}, 505)

	successes := success.AddSuccesses(map[any]int{
		"":           201,
		customType{}: 200,
	}, 505)

	r := gin.Default()

	r.GET("/string/:str", handlers.NewGinHandler[string, string](customErrors, successes).HandleParam("str", respondMessageString))
	r.GET("/str", handlers.NewGinHandler[string, string](customErrors, successes).HandleParam("string", errorMessageString))
	r.GET("/int", handlers.NewGinHandler[int, int](customErrors, successes).HandleQuery("integer", errorMessageInt))

	postHandlers := handlers.NewGinHandler[customType, customType](customErrors, successes)

	r.POST("/custom", postHandlers.HandleJSON(respondMessageCustomType))

	err := r.Run()
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
