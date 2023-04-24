package handlers

import (
	"gin-handler/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler[I, O any] struct {
	errorHandler   types.ErrorHandler
	successHandler types.SuccessHandler
}

func (h handler[I, O]) HandleJSON(fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		var input I

		err := c.BindJSON(&input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		output, err := fn(c, input)
		if err != nil {
			errorMessage, statusCode := h.errorHandler(err.Error())
			c.JSON(statusCode, errorMessage)
			return
		}

		successMessage, successesStatusCode := h.successHandler(output)
		c.JSON(successesStatusCode, successMessage)
		return
	}
}

func (h handler[I, O]) HandleParam(param string, fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		input := c.Param(param)

		converted, err := convert[I](input)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		output, err := fn(c, converted.(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler(err.Error())
			c.JSON(statusCode, errorMessage)
			return
		}

		successMessage, successesStatusCode := h.successHandler(output)
		c.JSON(successesStatusCode, successMessage)
		return
	}
}

func (h handler[I, O]) HandleQuery(query string, fn types.InputFunction[I, O]) func(*gin.Context) {
	return func(c *gin.Context) {
		input := c.Query(query)

		converted, err := convert[I](input)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		output, err := fn(c, converted.(I))
		if err != nil {
			errorMessage, statusCode := h.errorHandler(err.Error())
			c.JSON(statusCode, errorMessage)
			return
		}

		successMessage, successesStatusCode := h.successHandler(output)
		c.JSON(successesStatusCode, successMessage)
		return
	}
}

func Newhandler[I, O any](errorHandler types.ErrorHandler, successHandler types.SuccessHandler) types.Handler[I, O] {
	var handle handler[I, O]

	if errorHandler == nil {
		handle.errorHandler = returnError
	} else {
		handle.errorHandler = errorHandler
	}

	if successHandler == nil {
		handle.successHandler = returnSuccess
	} else {
		handle.successHandler = successHandler
	}

	return handle
}
