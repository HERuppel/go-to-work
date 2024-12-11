package routes

import (
	"go-to-work/internal/config"
	"go-to-work/internal/controllers"
	"go-to-work/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type ControllerContainer struct {
	UserController *controllers.UserController
	AuthController *controllers.AuthController
	JobController  *controllers.JobController
}

var router = gin.Default()

func Initialize(container *ControllerContainer) {
	getRoutes(container)

	router.Use(middlewares.CORS())

	router.Run(":" + config.Port)
}

func getRoutes(container *ControllerContainer) {
	v1 := router.Group("/v1")

	addAuthRoutes(v1, container.AuthController)
	addUserRoutes(v1, container.UserController)
	addJobRoutes(v1, container.JobController)
}
