package main

import (
	"context"
	"fmt"
	"gin-handler/handlers"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/string", handlers.NewHandler[string](messageString).Query("str").Response(respondMessageString).Handle)
	r.PUT("/si/:newString", handlers.NewHandler[int, string](messageIntToString).Param("newString").Handle)
	r.POST("/", handlers.NewHandler[tests](messageJSON).Handle)
	r.DELETE("/delete/:del", handlers.NewHandler[string](messageDelete).Param("del").Handle)
	err := r.Run()
	if err != nil {
		return
	}
}

func messageString(ctx context.Context, i string) (string, error) {
	return i, nil
}

func respondMessageString(c *gin.Context, resp string, err error) {
	if err != nil {
		return
	}
	c.JSONP(200, gin.H{resp: "Returned from custom response"})
}

func messageIntToString(ctx context.Context, i int) (string, error) {
	return strconv.Itoa(i), nil
}

type tests struct {
	Brusque bool `json:"brusque"`
}

func messageJSON(ctx context.Context, i tests) (tests, error) {
	return i, nil
}

func messageDelete(ctx context.Context, item string) (string, error) {
	return fmt.Sprintf("%v was deleted", item), nil
}
