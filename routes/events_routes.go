package routes

import (
	"example.com/rest-api-events/http_handlers"
	"example.com/rest-api-events/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(server *gin.Engine) {
	// call via curl: curl -s -X GET "http://localhost:8080/events" | json_pp
	server.GET("/events", http_handlers.GetEvents)
	server.GET("/events/:id", http_handlers.GetEventByID)
	server.POST("/events", middlewares.Authenticate, http_handlers.CreateEvent)
	server.PUT("/events/:id", http_handlers.UpdateEvent)
	server.DELETE("/events/:id", http_handlers.DeleteEvent)
}
