package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest_api_example.com/models"
	"rest_api_example.com/utils"
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

	context.JSON(http.StatusCreated, gin.H{"message": "user created!!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user data !!!!"})
		return
	}

	err = user.ValidateCred()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token !!!!"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "logined!!", "token": token})

}
