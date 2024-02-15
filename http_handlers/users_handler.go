package http_handlers

import (
	"example.com/rest-api-events/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to parse data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user!", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func Login(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to parse data", "error": err.Error()})
		return
	}

	isValid, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login the user!", "error": err.Error()})
		return
	}

	if !isValid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User unauthorized!", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged!"})
}
