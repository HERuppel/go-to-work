package routes

import (
	"go-to-work/internal/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	auth := rg.Group("/user")

	auth.GET("/:id", userController.GetUser)
}
