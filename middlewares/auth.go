package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest_api_example.com/utils"
)

func Authenticated(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Unauthorized"})

	}
	context.Set("userId", userId)
	context.Next()
}
