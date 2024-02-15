package main

import (
	"example.com/rest-api-events/db"
	"example.com/rest-api-events/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

// to install the third-party Gin package to project: go get -u github.com/gin-gonic/gin
// to study: https://www.jetbrains.com/guide/go/tutorials/rest_api_series/gin/

func main() {

	// initialize DB connection
	db.InitDB()

	server := gin.Default()

	// Register all endpoints and theirs handler function
	routes.RegisterRoutes(server)

	err := server.Run(":8080") // by default will be localhost:
	if err != nil {
		fmt.Println(err)
		return
	}
}
