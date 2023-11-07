package handlers

import (
	"github.com/gabrielcervante/handler/converter"
	customerrors "github.com/gabrielcervante/handler/custom_errors"
	"github.com/gabrielcervante/handler/success"
	"github.com/gabrielcervante/handler/types"
	"github.com/gabrielcervante/handler/utils"
	"github.com/gin-gonic/gin"
)

type ginHandler[I, O comparable] struct {
	errorHandler   customerrors.Errors
	successHandler success.Success
}

func (h ginHandler[I, O]) Handle(fn any) func(*gin.Context) {
	return func(c *gin.Context) {
		output, err := utils.FindFuncType[I, O](c, *new(I), fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (h ginHandler[I, O]) HandleJSON(fn any) func(*gin.Context) {
	return func(c *gin.Context) {
		var input I

		err := c.BindJSON(&input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		output, err := utils.FindFuncType[I, O](c, input, fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (h ginHandler[I, O]) HandleParam(param string, fn any) func(*gin.Context) {
	return func(c *gin.Context) {
		input := c.Param(param)

		converted, err := converter.Convert[O](input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		output, err := utils.FindFuncType[I, O](c, converted.(I), fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (h ginHandler[I, O]) HandleQuery(query string, fn any) func(*gin.Context) {
	return func(c *gin.Context) {
		input := c.Query(query)

		converted, err := converter.Convert[I](input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}
		output, err := utils.FindFuncType[I, O](c, converted.(I), fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		h.ginOutPut(c, output, successesStatusCode)
	}
}

func (ginHandler[I, O]) ginOutPut(ctx *gin.Context, output O, statusCode int) {
	if output == *new(O) {
		ctx.Status(statusCode)
	}

	ctx.JSON(statusCode, output)
}

func NewGinHandler[I, O comparable](errorHandler customerrors.Errors, successHandler success.Success) types.GinHandler[I, O] {
	return ginHandler[I, O]{errorHandler: errorHandler, successHandler: successHandler}
}
