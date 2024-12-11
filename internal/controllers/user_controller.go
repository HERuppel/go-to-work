package controllers

import (
	"go-to-work/internal/authentication"
	usecases "go-to-work/internal/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase *usecases.UserUseCase
}

func NewUserController(userUseCase *usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (userController *UserController) GetUser(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id, err := authentication.ExtractUserId(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_USER_ID"})
		return
	}

	user, err := userController.userUseCase.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_ERROR"})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "USER_NOT_FOUND"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
