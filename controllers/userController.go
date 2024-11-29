package controllers

import (
	"be-golang/connection"
	"be-golang/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateUserInput struct {
	Name     string `josn:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
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
	var input ValidateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
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
	var user models.User
	if err := connection.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//validate input
	var input ValidateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	//update user
	connection.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User Updated Successfully",
		"data":    user,
	})
}

func DeleteUser(c *gin.Context) {
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
