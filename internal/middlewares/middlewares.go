package middlewares

import (
	"go-to-work/internal/authentication"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "AUTHORIZATION_HEADER_REQUIRED"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "INVALID_AUTHORIZATION_HEADER"})
			ctx.Abort()
			return
		}

		if err := authentication.ValidateToken(tokenString); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
