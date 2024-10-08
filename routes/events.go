package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Fidel-wolee/event-booking/models"
)

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Events fetched successfully", "events": events})
}

func CreateEvent(c *gin.Context) {
    // Retrieve userId from the context
    userId, exists := c.Get("userId")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
        return
    }

    // Ensure that userId is an int64, as it's set in AuthMiddleware
    userIdInt64, ok := userId.(int64)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
        return
    }

    // Print the user ID for debugging purposes
    fmt.Printf("The user id is %v", userIdInt64)

    var event models.Event

    // Bind the request body to the event model
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
        return
    }

    // Assign the userId to the event
    event.UserID = userIdInt64

    // Save the event
    if err := event.Save(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
        return
    }

    // Success response
    c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}


func GetEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse event id."})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse request data."})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
