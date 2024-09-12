package routes

import (
	"net/http"
	"strconv"

	"github.com/Fidel-wolee/event-booking/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})

}

func getUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse event id."})
		return
	}
	user, _ := models.GetUser(userId)

	c.JSON(http.StatusOK, user)
}
