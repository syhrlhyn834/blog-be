package controllers

import (
	"be-golang/connection"
	"be-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Findusers(c *gin.Context) {
	var users []models.User
	connection.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data users",
		"data":    users,
	})
}
