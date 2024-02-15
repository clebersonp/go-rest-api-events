package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes - Routes all available endpoint and call their http_handlers from other file
func RegisterRoutes(server *gin.Engine) {
	RegisterEventsRoutes(server)
	RegisterUsersRoutes(server)
}
