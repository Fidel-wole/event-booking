package main

import (
	"github.com/Fidel-wolee/event-booking/db"
	"github.com/Fidel-wolee/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
    routes.RegisterRoutes(server)
	server.Run(":8080") // Run the server on port 8080
}
