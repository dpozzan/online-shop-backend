package routes

import (
	"net/http"

	"github.com/dpozzan/models"
	"github.com/gin-gonic/gin"
)

func registerNewUser(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request during Sign up"})
		return
	}

	id, err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register the new user"})
		return
	}

	user.ID = id

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func loginUser(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse body"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	

}