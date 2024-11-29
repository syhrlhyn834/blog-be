package controllers

import (
	"be-golang/connection"
	"be-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateCategoryInput struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

func FindCategories(c *gin.Context) {
	var categories []models.Category
	connection.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Category",
		"data":    categories,
	})
}
