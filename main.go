package main

import (
	"net/http"
	"strconv"

	"github.com/Fidel-wolee/event-booking/db"
	"github.com/Fidel-wolee/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)   // GET endpoint to fetch events
	server.POST("/event", createEvent) // POST endpoint to create a new event
    server.GET("/events/:id", getEvent)
	server.Run(":8080") // Run the server on port 8080
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Events fetched successfully", "events": events})
}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func getEvent(c *gin.Context){
  eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
  if err != nil{
	c.JSON(http.StatusBadRequest, gin.H{"message":"Cound not parse event id."})
  return
}

event, err := models.GetEventByID(eventId)
if err != nil{
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
   return
}
c.JSON(http.StatusOK, event)
}