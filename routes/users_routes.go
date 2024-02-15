package routes

import (
	"example.com/rest-api-events/http_handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(server *gin.Engine) {
	server.POST("/users", http_handlers.Signup)
	server.POST("/login", http_handlers.Login)
}
