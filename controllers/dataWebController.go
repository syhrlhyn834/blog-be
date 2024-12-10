package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateDataWebInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Favico      string `json:"favico" binding:"required"`
	Logo        string `json:"logo" binding:"required"`
	Footer      string `json:"footer" binding:"required"`
}

func FindDataWeb(c *gin.Context) {
	var dataWebs []models.Dataweb
	connection.DB.Find(&dataWebs)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Web",
		"data":    dataWebs,
	})
}

func StoreDataWeb(c *gin.Context) {
	middleware.AuthRequired()(c)
	var input ValidateDataWebInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	dataWebs := models.Dataweb{
		Title:       input.Title,
		Description: input.Description,
		Favico:      input.Favico,
		Logo:        input.Logo,
		Footer:      input.Footer,
	}
	connection.DB.Create(&dataWebs)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Web",
		"data":    dataWebs,
	})
}

func FindDataWebById(c *gin.Context) {
	var dataWebs models.Dataweb
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&dataWebs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail Data Web By ID : " + c.Param("id"),
		"data":    dataWebs,
	})
}

func UpdateDataWeb(c *gin.Context) {
	middleware.AuthRequired()(c)
	var dataWebs models.Dataweb
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&dataWebs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input ValidateDataWebInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	connection.DB.Model(&dataWebs).Updates(models.Dataweb{
		Title:       input.Title,
		Description: input.Description,
		Favico:      input.Favico,
		Logo:        input.Logo,
		Footer:      input.Footer,
	})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Web Updated Successfully",
		"data":    dataWebs,
	})
}

func DeleteDataWeb(c *gin.Context) {
	middleware.AuthRequired()(c)
	var dataWebs models.Dataweb
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&dataWebs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	connection.DB.Delete(&dataWebs)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Web Deleted Successfully",
	})
}
