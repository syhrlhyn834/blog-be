package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type ValidatePostInput struct {
	Title       string `json:"title" binding:"required,max=255"`
	Slug        string `json:"slug" binding:"required,max=255"`
	CategoryID  int    `json:"category_id" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Image       string `json:"image" binding:"omitempty,url"`
	Status      string `json:"status" binding:"required,oneof=draft published archive"`
}

func FindPost(c *gin.Context) {
	var posts []models.Post
	if err := connection.DB.Preload("Category").Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data posts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List data Post",
		"data":    posts,
	})
}

func StorePost(c *gin.Context) {
	middleware.AuthRequired()(c)

	// Validasi input JSON untuk data selain gambar
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	// Validasi apakah kategori ada di database
	var category models.Category
	if err := connection.DB.First(&category, input.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category not found",
		})
		return
	}

	// Proses upload file gambar
	file, err := c.FormFile("image")
	var imagePath string
	if err == nil {
		// Jika file gambar ditemukan, simpan file ke folder `image`
		uploadPath := "./src/image"
		imagePath = filepath.Join(uploadPath, file.Filename)

		// Simpan file ke folder
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}
	}

	generatedSlug := slug.Make(input.Title)
	// Simpan data post ke database
	post := models.Post{
		Title:       input.Title,
		Slug:        generatedSlug,
		CategoryID:  input.CategoryID,
		UserID:      input.UserID,
		Description: input.Description,
		Content:     input.Content,
		Image:       imagePath, // Path file yang disimpan
		Status:      input.Status,
	}

	if err := connection.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post berhasil dibuat",
		"data":    post,
	})
}
