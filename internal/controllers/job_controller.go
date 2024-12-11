package controllers

import (
	"go-to-work/internal/authentication"
	"go-to-work/internal/models"
	usecases "go-to-work/internal/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	jobUseCase *usecases.JobUseCase
}

func NewJobController(jobUseCase *usecases.JobUseCase) *JobController {
	return &JobController{
		jobUseCase: jobUseCase,
	}
}

func (jobController *JobController) Create(ctx *gin.Context) {
	var job models.Job

	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	recruiterId, err := authentication.ExtractUserId(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_USER_ID"})
		return
	}

	job.RecruiterId = recruiterId

	createdId, err := jobController.jobUseCase.Create(ctx, job)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": createdId})
}
