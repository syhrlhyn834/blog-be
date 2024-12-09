package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateMenuInput struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

func FindMenu(c *gin.Context) {
	var menus []models.Menu
	connection.DB.Find(&menus)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Menu",
		"data":    menus,
	})
}

func StoreMenu(c *gin.Context) {
	middleware.AuthRequired()(c)
	var input ValidateMenuInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	menus := models.Menu{
		Name: input.Name,
		Url:  input.Url,
	}
	connection.DB.Create(&menus)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Menu Created",
		"data":    menus,
	})
}

func FindMenuById(c *gin.Context) {
	var menus models.Menu
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&menus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail Data Menu By ID : " + c.Param("id"),
		"data":    menus,
	})
}

func UpdateMenu(c *gin.Context) {
	middleware.AuthRequired()(c)
	var menus models.Menu
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&menus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input ValidateMenuInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	connection.DB.Model(&menus).Updates(models.Menu{
		Name: input.Name,
		Url:  input.Url,
	})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Menu Updated Successfully",
		"data":    menus,
	})
}

func DeleteMenu(c *gin.Context) {
	middleware.AuthRequired()(c)
	var menus models.Menu
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&menus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	connection.DB.Delete(&menus)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Menu Deleted Successfully",
	})
}
