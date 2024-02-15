package main

import (
	"example.com/rest-api-events/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// to install the third-party Gin package to project: go get -u github.com/gin-gonic/gin

func main() {

	server := gin.Default()

	// call via curl: curl -s -X GET "http://localhost:8080/events" | json_pp
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080") // by default will be localhost:
	if err != nil {
		fmt.Println(err)
		return
	}
}

// getEvents - will be used as named function by handler
func getEvents(context *gin.Context) {
	// with gin all we need to do we do with context
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(context *gin.Context) {
	event := models.Event{}
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Failed to parse data: %v\n", err)})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
