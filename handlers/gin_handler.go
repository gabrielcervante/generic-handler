package handlers

import (
	"github.com/gabrielcervante/handler/converter"
	customerrors "github.com/gabrielcervante/handler/custom_errors"
	"github.com/gabrielcervante/handler/success"
	"github.com/gabrielcervante/handler/types"
	"github.com/gabrielcervante/handler/utils"
	"github.com/gin-gonic/gin"
)

type ginHandler[I, O any] struct {
	errorHandler   customerrors.Errors
	successHandler success.Success
}

func (h ginHandler[I, O]) Handle(fn any) func(*gin.Context) {
	return func(c *gin.Context) {
		var input I

		output, err := utils.FindFuncType[I, O](c, input, fn)
		if err != nil {
			errorMessage, statusCode := h.errorHandler.ReturnError(err)
			c.JSON(statusCode, errorMessage)
			return
		}

		successesStatusCode := h.successHandler.ReturnSuccess(*new(I))
		c.JSON(successesStatusCode, output)
		return
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
		c.JSON(successesStatusCode, output)
		return
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
		c.JSON(successesStatusCode, output)
		return
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
		c.JSON(successesStatusCode, output)
		return
	}
}

func NewGinHandler[I, O any](errorHandler customerrors.Errors, successHandler success.Success) types.GinHandler[I, O] {
	return ginHandler[I, O]{errorHandler: errorHandler, successHandler: successHandler}
}
