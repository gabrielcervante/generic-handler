package types

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type InputFunction[I, O any] func(context.Context, I) (O, error)

type OutputFunction[O any] func(context.Context) (O, error)

type ErrorFunction[I any] func(context.Context, I) error

type GinHandler[I, O any] interface {
	HandleJSON(any) func(*gin.Context)
	HandleParam(string, any) func(*gin.Context)
	HandleQuery(string, any) func(*gin.Context)
}

type FiberHandler[I, O any] interface {
	HandleJSON(any) func(*fiber.Ctx) error
	HandleParam(string, any) func(*fiber.Ctx) error
	HandleQuery(string, any) func(*fiber.Ctx) error
}
