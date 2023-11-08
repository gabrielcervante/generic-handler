package types

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type Convert[O any] func(input string, converter any) (any, error)

type InputFunction[I, O any] func(context.Context, I) (O, error)

type OutputFunction[O any] func(context.Context) (O, error)

type ErrorFunction[I any] func(context.Context, I) error

type GinHandler[Converter, I, O comparable] interface {
	Handle(InputFunction[I, O]) func(*gin.Context)
	HandleJSON(InputFunction[I, O]) func(*gin.Context)
	HandleParam(string, InputFunction[I, O]) func(*gin.Context)
	HandleQuery(string, InputFunction[I, O]) func(*gin.Context)
}

type FiberHandler[Converter, I, O any] interface {
	Handle(InputFunction[I, O]) func(*fiber.Ctx) error
	HandleJSON(InputFunction[I, O]) func(*fiber.Ctx) error
	HandleParam(string, InputFunction[I, O]) func(*fiber.Ctx) error
	HandleQuery(string, InputFunction[I, O]) func(*fiber.Ctx) error
}
