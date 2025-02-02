package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_rest_api/utils"
)

func Authorization(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is required"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return

	}

	context.Set("userId", userId)
	context.Next()

}
