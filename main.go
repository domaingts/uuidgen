package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()
	router.GET("/uuidgen/:version", func(ctx *gin.Context) {
		version, ok := ctx.Params.Get("version")
		if !ok || version == "" {
			ctx.String(http.StatusBadRequest, "Bad Param")
			return
		}
		id, err := uuid.NewV7()
		if err != nil {
			ctx.String(http.StatusBadGateway, err.Error())
			return
		}
		ctx.String(http.StatusOK, id.String())
	})
	err := router.Run(":13000")
	if err != nil {
		panic(err)
	}
}
