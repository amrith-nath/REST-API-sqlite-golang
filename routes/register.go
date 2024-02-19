package routes

import (
	"net/http"
	"strconv"

	"example.com/booking-app/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("user_id")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Register event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event register successfully"})
}

func cancelForEvent(context *gin.Context) {
	userId := context.GetInt64("user_id")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Cancel(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Cancel event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Cancelled successfully"})
}
