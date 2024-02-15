package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes - Routes all available endpoint and call their handlers from other file
func RegisterRoutes(server *gin.Engine) {
	// call via curl: curl -s -X GET "http://localhost:8080/events" | json_pp
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)

}
