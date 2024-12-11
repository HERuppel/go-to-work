package routes

import (
	"go-to-work/internal/controllers"
	"go-to-work/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func addJobRoutes(rg *gin.RouterGroup, jobController *controllers.JobController) {
	user := rg.Group("/job")
	user.Use(middlewares.Authenticate())

	user.POST("/", jobController.Create)
}
