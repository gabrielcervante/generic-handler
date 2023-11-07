package utils

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gabrielcervante/handler/types"
)

type withoutOutput[I any] func(ctx context.Context, i I, fn types.ErrorFunction[I]) error

func FindFuncType[I, O any](ctx context.Context, input I, fnToFindType any) (O, error) {
	if reflect.TypeOf(fnToFindType).NumIn() == 2 && reflect.TypeOf(fnToFindType).NumOut() == 2 {
		return inputFunction[I, O](ctx, input, fnToFindType.(types.InputFunction[I, O]))
	}

	if reflect.TypeOf(fnToFindType).NumIn() == 1 && reflect.TypeOf(fnToFindType).NumOut() == 2 {
		return outputFunction[O](ctx, fnToFindType.(types.OutputFunction[O]))
	}

	if reflect.TypeOf(fnToFindType).NumIn() == 2 && reflect.TypeOf(fnToFindType).NumOut() == 1 {
		return emptyOutput[I, O](ctx, input, fnToFindType.(types.ErrorFunction[I]), errorFunction[I])
	}

	panic("Sorry, you provided a wrong type " + fmt.Sprintf("%T", fnToFindType))
}

func inputFunction[I, O any](ctx context.Context, i I, fn types.InputFunction[I, O]) (O, error) {
	return fn(ctx, i)
}

func outputFunction[O any](ctx context.Context, fn types.OutputFunction[O]) (O, error) {
	return fn(ctx)
}

func emptyOutput[I, O any](ctx context.Context, i I, errFn types.ErrorFunction[I], fn withoutOutput[I]) (_ O, err error) {
	err = fn(ctx, i, errFn)
	return
}

func errorFunction[I any](ctx context.Context, i I, fn types.ErrorFunction[I]) error {
	return fn(ctx, i)
}
