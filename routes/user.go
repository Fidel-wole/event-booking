package routes

import (
	"net/http"

	"github.com/Fidel-wolee/event-booking/models"
	"github.com/Fidel-wolee/event-booking/utils"
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
	userId, exists := c.Get("userId")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
        return
    }
    userId, ok := userId.(int64)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user ID"})
        return
    }
    userIdInt64, ok := userId.(int64)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user ID type"})
        return
    }
    user, err := models.GetUser(userIdInt64)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve user", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func login(c *gin.Context) {
    var user models.LoginUser
    err := c.ShouldBindJSON(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user data."})
        return
    }

    // Validate credentials and get the user with ID
    validatedUser, err := user.ValidateCredentials()
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "Invalid credentials.",
            "error":   err.Error(),
        })
        return
    }

    // Generate token with the validated user's email and ID
    token, err := utils.GenerateToken(validatedUser.Email, validatedUser.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
