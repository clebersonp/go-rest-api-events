package http_handlers

import (
	"example.com/rest-api-events/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// In this file we create all handler function for endpoint registered into routes.go file
func convertToInt64(context *gin.Context, parameterName string) (num int64, err error) {
	return strconv.ParseInt(context.Param(parameterName), 10, 64)
}

// GetEvents - will be used as named function by handler
func GetEvents(context *gin.Context) {
	// with gin all we need to do we do with context
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve data", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	event := models.Event{}
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse data", "error": err.Error()})
		return
	}

	// set the userId that is requiring to save this event from this context
	// get userId from this context keys
	event.UserID = context.GetInt64("user_id")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event!", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func GetEventByID(context *gin.Context) {
	id, err := convertToInt64(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID can't converted to integer!", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event by ID!", "error": err.Error()})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event Not Found!", "error": err})
		return
	}

	context.JSON(http.StatusOK, event)
}

func UpdateEvent(context *gin.Context) {
	id, err := convertToInt64(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID can't converted to integer!", "error": err.Error()})
		return
	}

	eventDb, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get eventDb by ID!", "error": err.Error()})
		return
	}
	if eventDb == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event Not Found!", "error": err})
		return
	}

	updatedEvent := models.Event{}
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse data", "error": err.Error()})
		return
	}

	updatedEvent.ID = id

	// get userId from this context keys from auth middleware
	userId := context.GetInt64("user_id")
	if eventDb.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to update this event", "error": nil})
		return
	}

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func DeleteEvent(context *gin.Context) {
	id, err := convertToInt64(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID can't converted to integer!", "error": err.Error()})
		return
	}

	eventDb, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get eventDb by ID!", "error": err.Error()})
		return
	}
	if eventDb == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event Not Found!", "error": err})
		return
	}

	// get userId from this context keys from auth middleware
	userId := context.GetInt64("user_id")
	if eventDb.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to delete this event", "error": nil})
		return
	}

	err = eventDb.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event!", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
