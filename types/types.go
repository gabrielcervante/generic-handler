package types

import (
	"context"

	"github.com/gin-gonic/gin"
)

type InputFunction[I, O any] func(context.Context, I) (O, error)

type OutputFunction[O any] func(context.Context) (O, error)

type ErrorHandler func(string) (string, int)

type SuccessHandler func(any) (any, int)

type Handler[I, O any] interface {
	HandleJSON(InputFunction[I, O]) func(*gin.Context)
	HandleParam(string, InputFunction[I, O]) func(*gin.Context)
	HandleQuery(string, InputFunction[I, O]) func(*gin.Context)
}
