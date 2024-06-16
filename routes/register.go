package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest_api_example.com/models"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("user_id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id !!!!"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found !!!!"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event !!!!"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!!"})

}

func cancelEvent(context *gin.Context) {
	userId := context.GetInt64("user_id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId
	err = event.CancelReg(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel reg!!!!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "canceled"})

}
