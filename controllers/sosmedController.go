package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateSosmedInput struct {
	Name     string `json:"name" binding:"required"`
	Logo     string `json:"logo" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func FindSosmed(c *gin.Context) {
	var sosmed []models.Sosmed
	connection.DB.Find(&sosmed)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    sosmed,
	})
}

func StoreSosmed(c *gin.Context) {
	middleware.AuthRequired()(c)
	var input ValidateSosmedInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}
	sosmed := models.Sosmed{
		Name:     input.Name,
		Logo:     input.Logo,
		Url:      input.Url,
		Username: input.Username,
	}
	connection.DB.Create(&sosmed)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sosmed Created",
		"data":    sosmed,
	})
}

func FindSosmedById(c *gin.Context) {
	var sosmed models.Sosmed
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail Data Sosmed By ID : " + c.Param("id"),
		"data":    sosmed,
	})
}

func UpdateSosmed(c *gin.Context) {
	middleware.AuthRequired()(c)
	var sosmed models.Sosmed
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input ValidateSosmedInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}
	connection.DB.Model(&sosmed).Updates(models.Sosmed{
		Name:     input.Name,
		Logo:     input.Logo,
		Url:      input.Url,
		Username: input.Username,
	})
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sosmed Updated Successfully",
		"data":    sosmed,
	})
}

func DeleteSosmed(c *gin.Context) {
	middleware.AuthRequired()(c)
	var sosmed models.Sosmed
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	connection.DB.Delete(&sosmed)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sosmed Deleted Successfully",
	})
}
