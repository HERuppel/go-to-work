package routes

import (
	"go-to-work/internal/config"
	"go-to-work/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ControllerContainer struct {
	UserController *controllers.UserController
}

var router = gin.Default()

func Initialize(container *ControllerContainer) {
	getRoutes(container)
	router.Run(":" + config.Port)
}

func getRoutes(container *ControllerContainer) {
	v1 := router.Group("/v1")

	addUserRoutes(v1, container.UserController)
}
