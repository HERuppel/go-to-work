package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "AUTH")
	})
}
