package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type ValidateCategoryInput struct {
	Name string `json:"name" binding:"required"`
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

func StoreCategory(c *gin.Context) {
	middleware.AuthRequired()(c)
	var input ValidateCategoryInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	generatedSlug := slug.Make(input.Name)

	categories := models.Category{
		Name: input.Name,
		Slug: generatedSlug,
	}
	connection.DB.Create(&categories)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category Created",
		"data":    categories,
	})
}

func FindCategoryById(c *gin.Context) {
	var categories models.Category
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail Data Category By ID : " + c.Param("id"),
		"data":    categories,
	})
}

func UpdateCategory(c *gin.Context) {
	middleware.AuthRequired()(c)
	var categories models.Category
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input ValidateCategoryInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors})
			return
		}
	}
	generatedSlug := slug.Make(input.Name)

	connection.DB.Model(&categories).Updates(models.Category{
		Name: input.Name,
		Slug: generatedSlug,
	})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category Updated Successfully",
		"data":    categories,
	})
}

func DeleteCategory(c *gin.Context) {
	middleware.AuthRequired()(c)
	var categories models.Category
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	connection.DB.Delete(&categories)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category Deleted Successfully",
		"data":    categories,
	})
}
