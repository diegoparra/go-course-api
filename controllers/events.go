package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-course/api/models"
)

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"Message": "Could not fetch the db data"},
		)
	}
	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"Message": "Could not parse the userId"},
		)
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"Message": "Could not fetch the event from db"},
		)
	}

	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.UserId = 1
	event.Save()
	c.JSON(http.StatusCreated, event)
}

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"Message": "Could not parse the eventId"},
		)
		return
	}

	var event models.Event

	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"Message": err.Error()},
		)
		return
	}

	err = models.UpdateEvent(event, eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"Message": "Could not update event"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Event updated successfully"})
}

func DeleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"Message": "Could not parse the eventId"},
		)
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"Message": "Could not delete the event"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Event deleted successfully"})
}
