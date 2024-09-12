package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)   // GET endpoint to fetch events
	server.POST("/event", CreateEvent) // POST endpoint to create a new event
	server.GET("/events/:id", GetEvent)
	server.PUT("/event/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)
}
