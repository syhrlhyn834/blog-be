package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidateUserInput struct {
	Name     string `josn:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Findusers(c *gin.Context) {
	var users []models.User
	connection.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data User",
		"data":    users,
	})
}

func StoreUser(c *gin.Context) {
	middleware.AuthRequired()(c)
	var input ValidateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	connection.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User Data created.",
		"data":    user,
	})
}

func FindUserById(c *gin.Context) {
	var user models.User
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail Data user By ID : " + c.Param("id"),
		"data":    user,
	})
}

func UpdateUser(c *gin.Context) {
	middleware.AuthRequired()(c)
	var user models.User
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validasi input
	var input ValidateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	// Update data user
	connection.DB.Model(&user).Updates(input)

	// Kirim respon sukses
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User Updated Successfully",
		"data":    user,
	})
}

func DeleteUser(c *gin.Context) {
	middleware.AuthRequired()(c)
	var user models.User
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//delete user
	connection.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User Deleted Successfully",
	})
}
