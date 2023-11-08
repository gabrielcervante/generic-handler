package example

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/gabrielcervante/handler/adapter"
	"github.com/gabrielcervante/handler/converter"
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

	convert := converter.Converter[int]{ConvertWithError: strconv.Atoi}

	r := fiber.New()

	r.Get("/string/:str", handlers.NewFiberHandler[string, string, string](customErrors, successes).HandleParam("str", respondMessageString))
	r.Get("/str", handlers.NewFiberHandler[string, string, string](customErrors, successes).HandleParam("string", errorMessageString))
	r.Get("/int", handlers.NewFiberHandler[int, int, int](customErrors, successes, convert).HandleQuery("integer", errorMessageInt))
	r.Get("/int", handlers.NewFiberHandler[string, struct{}, string](customErrors, successes).Handle(adapter.NoInput[struct{}, string](testNoInput)))
	r.Get("/int", handlers.NewFiberHandler[string, string, struct{}](customErrors, successes).HandleQuery("integer", adapter.NoOutput[string, struct{}](testNoOutput)))

	postHandlers := handlers.NewFiberHandler[customType, customType, customType](customErrors, successes, converter.Converter[customType]{})

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

	convert := converter.Converter[int]{ConvertWithError: strconv.Atoi}

	r := gin.Default()

	r.GET("/string/:str", handlers.NewGinHandler[string, string, string](customErrors, successes).HandleParam("str", respondMessageString))
	r.GET("/str", handlers.NewGinHandler[string, string, string](customErrors, successes).HandleParam("string", errorMessageString))
	r.GET("/int", handlers.NewGinHandler[int, int, int](customErrors, successes, convert).HandleQuery("integer", errorMessageInt))
	r.GET("/int", handlers.NewGinHandler[struct{}, struct{}, string](customErrors, successes).Handle(adapter.NoInput[struct{}, string](testNoInput)))
	r.GET("/int", handlers.NewGinHandler[string, string, struct{}](customErrors, successes).HandleQuery("integer", adapter.NoOutput[string, struct{}](testNoOutput)))

	postHandlers := handlers.NewGinHandler[customType, customType, customType](customErrors, successes)

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

func testNoInput(ctx context.Context) (string, error) {
	return "Nothing here", nil
}

func testNoOutput(ctx context.Context, input string) error {
	fmt.Println(input)
	return nil
}
