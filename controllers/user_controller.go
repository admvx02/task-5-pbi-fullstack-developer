package controllers

import (
	"app/helpers"
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helpers.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters"})
		return
	}

	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword
	models.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}
