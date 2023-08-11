package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := helpers.GetUserIDFromToken(c)

	photo.UserID = userId
	models.DB.Create(&photo)

	c.JSON(http.StatusOK, photo)
}
