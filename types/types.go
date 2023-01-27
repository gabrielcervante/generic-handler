package types

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Operation[I, O any] func(context.Context, I) (O, error)
type Response[O any] func(*gin.Context, O, error)
