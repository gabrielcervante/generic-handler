package handlers

import (
	"github.com/gabrielcervante/handler/converter"
	customerrors "github.com/gabrielcervante/handler/custom_errors"
	"github.com/gabrielcervante/handler/success"
	"github.com/gabrielcervante/handler/types"
	"github.com/gin-gonic/gin"
)

type ginHandler[Converter, I, O comparable] struct {
	errorHandler   customerrors.Errors
	successHandler success.Success
	convert        converter.Converter[Converter]
}

func (h ginHandler[Converter, I, O]) Handle(fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		output, err := fn(c, *new(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (h ginHandler[Converter, I, O]) HandleJSON(fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		var input I

		err := c.BindJSON(&input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		output, err := fn(c, input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (h ginHandler[Converter, I, O]) HandleParam(param string, fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		input := c.Param(param)

		converted, err := h.convert.Convert(input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		output, err := fn(c, converted.(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (h ginHandler[Converter, I, O]) HandleQuery(query string, fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		input := c.Query(query)

		converted, err := h.convert.Convert(input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}
		output, err := fn(c, converted.(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (ginHandler[Converter, I, O]) ginOutPut(ctx *gin.Context, output O, statusCode int) {
	if output == *new(O) {
		ctx.Status(statusCode)
	}

	ctx.JSON(statusCode, output)
}

func NewGinHandler[Converter, I, O comparable](errorHandler customerrors.Errors, successHandler success.Success, convert ...converter.Converter[Converter]) types.GinHandler[Converter, I, O] {
	if convert != nil {
		return ginHandler[Converter, I, O]{errorHandler: errorHandler, successHandler: successHandler, convert: convert[0]}
	}
	return ginHandler[Converter, I, O]{errorHandler: errorHandler, successHandler: successHandler, convert: converter.Converter[Converter]{}}
}
