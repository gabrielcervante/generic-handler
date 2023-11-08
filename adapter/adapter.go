package adapter

import (
	"context"

	"github.com/gabrielcervante/handler/types"
)

func NoOutput[I, O any](fn types.ErrorFunction[I]) types.InputFunction[I, O] {
	return func(ctx context.Context, i I) (_ O, err error) {
		err = fn(ctx, i)
		return
	}
}

func NoInput[I, O any](fn types.OutputFunction[O]) types.InputFunction[I, O] {
	return func(ctx context.Context, _ I) (O, error) {
		return fn(ctx)
	}
}
