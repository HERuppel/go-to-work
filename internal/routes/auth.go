package routes

import (
	"go-to-work/internal/controllers"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup, authController *controllers.AuthController) {
	auth := rg.Group("/auth")

	auth.POST("/signup", authController.SignUp)
}
