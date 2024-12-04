package routes

import (
	"go-to-work/internal/controllers"
	"go-to-work/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	user := rg.Group("/user")
	user.Use(middlewares.Authenticate())

	user.GET("/", userController.GetUser)
}
