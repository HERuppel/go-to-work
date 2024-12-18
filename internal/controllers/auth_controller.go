package controllers

import (
	"go-to-work/internal/models"
	usecases "go-to-work/internal/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase *usecases.AuthUseCase
}

func NewAuthController(authUseCase *usecases.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

func (authController *AuthController) SignUp(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := authController.authUseCase.SignUp(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (authController *AuthController) SignIn(ctx *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_PAYLOAD"})
		return
	}

	signInResponse, err := authController.authUseCase.SignIn(ctx, credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": signInResponse.User, "authToken": signInResponse.AuthToken})
}

func (authController *AuthController) ConfirmAccount(ctx *gin.Context) {
	var payload struct {
		Email   string `json:"email" binding:"required,email"`
		PinCode string `json:"pin_code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_PAYLOAD"})
		return
	}

	err := authController.authUseCase.ConfirmAccount(ctx, payload.Email, payload.PinCode)
	if err != nil {
		if err.Error() == "INTERNAL_ERROR" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ACCOUNT_CONFIRM_WITH_SUCCESS"})
}
