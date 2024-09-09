package main

import (
	"net/http"

	"github.com/Fidel-wolee/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
    
	server.POST("/event", createEvent)
	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"message":"Events fetched sucessfully", "events":events})
}

func createEvent(context *gin.Context){
  var event models.Event
  err := context.ShouldBindJSON(&event)
  if err != nil{
	context.JSON(http.StatusBadRequest, gin.H{"message":"Could not padd request data"})
  }

  event.ID = 1
  event.UserID = 1

  event.Save()
  
  context.JSON(http.StatusCreated, gin.H{"message":"Event created", "event":event})
}