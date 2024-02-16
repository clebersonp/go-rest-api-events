package routes

import (
	"example.com/rest-api-events/http_handlers"
	"example.com/rest-api-events/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(server *gin.Engine) {
	// call via curl: curl -s -X GET "http://localhost:8080/events" | json_pp
	// example of endpoints without authentication
	server.GET("/events", http_handlers.GetEvents)
	server.GET("/events/:id", http_handlers.GetEventByID)

	// create a group to register the middlewares only once to authenticate user
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", middlewares.Authenticate, http_handlers.CreateEvent)
	authenticated.PUT("/events/:id", http_handlers.UpdateEvent)
	authenticated.DELETE("/events/:id", http_handlers.DeleteEvent)
	authenticated.POST("/events/:id/register", http_handlers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", http_handlers.CancelRegistration)
}
