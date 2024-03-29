package handlers

import (
	"github.com/gabrielcervante/handler/converter"
	customerrors "github.com/gabrielcervante/handler/custom_errors"
	"github.com/gabrielcervante/handler/success"
	"github.com/gabrielcervante/handler/types"
	"github.com/gofiber/fiber/v2"
)

type fiberHandler[I, O comparable] struct {
	errorHandler   customerrors.Errors
	successHandler success.Success
	convert        converter.Converter[I]
}

func (h fiberHandler[I, O]) Handle(fn types.InputFunction[I, O]) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		output, err := fn(c.Context(), *new(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (h fiberHandler[I, O]) HandleJSON(fn types.InputFunction[I, O]) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var input I

		err := c.BodyParser(&input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		output, err := fn(c.Context(), input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (h fiberHandler[I, O]) HandleParam(param string, fn types.InputFunction[I, O]) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		input := c.Params(param)

		converted, err := h.convert.Convert(input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		output, err := fn(c.Context(), converted.(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		return h.fiberOutPut(c, output, successesStatusCode)
	}
}

func (h fiberHandler[I, O]) HandleQuery(query string, fn types.InputFunction[I, O]) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		input := c.Query(query)

		converted, err := h.convert.Convert(input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			return c.Status(statusCode).JSON(errorMessage)
		}
		output, err := fn(c.Context(), converted.(I))
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

func NewFiberHandler[I, O comparable](errorHandler customerrors.Errors, successHandler success.Success, convert ...converter.Converter[I]) types.FiberHandler[I, O] {
	if convert != nil {
		return fiberHandler[I, O]{errorHandler: errorHandler, successHandler: successHandler, convert: convert[0]}
	}
	return fiberHandler[I, O]{errorHandler: errorHandler, successHandler: successHandler, convert: converter.Converter[I]{}}
}
