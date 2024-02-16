package middlewares

import (
	"example.com/rest-api-events/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(context *gin.Context) {
	authorization := context.Request.Header.Get("Authorization")
	token, found := strings.CutPrefix(authorization, "Bearer ")

	if authorization == "" || !found {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// set to context the user_id as key-value
	context.Set("user_id", userId)
	context.Next()
}
