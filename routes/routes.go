package routes

import (
	"github.com/gin-gonic/gin"
   "github.com/Fidel-wolee/event-booking/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	// Protected routes
	authorized := server.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/events", CreateEvent)
		authorized.PUT("/events/:id", UpdateEvent)
		authorized.DELETE("/events/:id", DeleteEvent)
		authorized.GET("/user/:id", getUser)
	}
}
