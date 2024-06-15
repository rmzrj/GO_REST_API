package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest_api_example.com/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user data !!!!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user data !!!!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"user created!!"})
}
