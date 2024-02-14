package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// to install the third-party Gin package to project: go get -u github.com/gin-gonic/gin

func main() {

	server := gin.Default()

	// call via curl: curl -s -X GET "http://localhost:8080/events" | json_pp
	server.GET("/events", getEvents)

	err := server.Run(":8080") // by default will be localhost:
	if err != nil {
		fmt.Println(err)
		return
	}
}

// getEvents - will be used as named function by handler
func getEvents(context *gin.Context) {
	// with gin all we need to do we do with context
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
