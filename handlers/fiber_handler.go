package handlers

import (
	"github.com/gabrielcervante/handler/converter"
	customerrors "github.com/gabrielcervante/handler/custom_errors"
	"github.com/gabrielcervante/handler/success"
	"github.com/gabrielcervante/handler/types"
	"github.com/gabrielcervante/handler/utils"
	"github.com/gofiber/fiber/v2"
)

type fiberHandler[I, O comparable] struct {
	errorHandler   customerrors.Errors
	successHandler success.Success
}

func (h fiberHandler[I, O]) Handle(fn any) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var input I

		output, err := utils.FindFuncType[I, O](c.Context(), input, fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (h fiberHandler[I, O]) HandleJSON(fn any) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var input I

		err := c.BodyParser(&input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		output, err := utils.FindFuncType[I, O](c.Context(), input, fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (h fiberHandler[I, O]) HandleParam(param string, fn any) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		input := c.Params(param)

		converted, err := converter.Convert[I](input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		output, err := utils.FindFuncType[I, O](c.Context(), converted.(I), fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (h fiberHandler[I, O]) HandleQuery(query string, fn any) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		input := c.Query(query)

		converted, err := converter.Convert[I](input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}
		output, err := utils.FindFuncType[I, O](c.Context(), converted.(I), fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (fiberHandler[I, O]) fiberOutPut(ctx *fiber.Ctx, output O, statusCode int) error {
	if output == *new(O) {
		return ctx.SendStatus(statusCode)
	}

	return ctx.Status(statusCode).JSON(output)
}

func NewFiberHandler[I, O comparable](errorHandler customerrors.Errors, successHandler success.Success) types.FiberHandler[I, O] {
	return fiberHandler[I, O]{errorHandler: errorHandler, successHandler: successHandler}
}
