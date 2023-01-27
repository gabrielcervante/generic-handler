package handlers

import (
	"fmt"
	"gin-handler/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler[I, O any] struct {
	newParam any
	newQuery any
	op       types.Operation[I, O]
	resp     types.Response[O]
}

func NewHandler[I, O any](handle types.Operation[I, O]) *Handler[I, O] {
	return &Handler[I, O]{
		op: handle,
	}
}

func (h *Handler[I, O]) Param(newParam string) *Handler[I, O] {
	h.newParam = newParam

	return h
}

func (h *Handler[I, O]) Query(newQuery string) *Handler[I, O] {
	h.newQuery = newQuery

	return h
}

func (h *Handler[I, O]) Response(response types.Response[O]) *Handler[I, O] {
	h.resp = response

	return h
}

func (h *Handler[I, O]) Handle(c *gin.Context) {
	var a O

	input, err := h.parse(c)
	if err != nil {
		h.response(c, a, err)
	}

	resp, err := h.op(c, input)
	h.response(c, resp, err)

	return
}

func (h *Handler[I, O]) parse(c *gin.Context) (I, error) {
	if h.newParam != nil {
		return h.getParam(c)
	} else if h.newQuery != nil {
		return h.getQuery(c)
	} else {
		return h.getJSON(c)
	}
}

func (h *Handler[I, O]) getParam(c *gin.Context) (I, error) {
	param := c.Param(h.newParam.(string))
	correctType, err := h.correctType(param)
	return correctType.(I), err
}

func (h *Handler[I, O]) getQuery(c *gin.Context) (I, error) {
	query := c.Query(h.newQuery.(string))
	correctType, err := h.correctType(query)
	return correctType.(I), err
}

func (h *Handler[I, O]) getJSON(c *gin.Context) (I, error) {
	var input I
	err := c.BindJSON(&input)
	return input, err
}

func (h *Handler[I, O]) correctType(value string) (returnedValue any, err error) {
	var genericType I

	switch fmt.Sprintf("%T", genericType) {
	case "int":
		returnedValue, err = strconv.Atoi(value)
		return
	default:
		return value, nil
	}
}

func (h *Handler[I, O]) response(c *gin.Context, resp O, err error) {
	if h.resp != nil {
		h.resp(c, resp, err)
	} else {
		if err != nil {
			c.JSONP(500, err)
			return
		}
		c.JSONP(200, resp)
	}
}
