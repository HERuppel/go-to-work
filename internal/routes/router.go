package routes

import (
	"go-to-work/internal/config"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Initialize() {
	getRoutes()
	router.Run(":" + config.Port)
}

func getRoutes() {
	v1 := router.Group("/v1")

	addAuthRoutes(v1)
}
