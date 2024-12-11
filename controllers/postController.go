package controllers

import (
	"be-golang/connection"
	"be-golang/middleware"
	"be-golang/models"
	"be-golang/resources"
	"fmt"
	"net/http"
	"os"
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
	Image       string `json:"image" binding:"required"`
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

	// Menangani Upload Gambar
	// Mendapatkan file gambar dari form-data
	file, err := c.FormFile("image") // "image" adalah nama form input di frontend
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Tentukan folder tujuan untuk menyimpan gambar
	imageFolder := "src/images"
	// Pastikan folder tujuan ada
	err = os.MkdirAll(imageFolder, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create image folder"})
		return
	}

	// Tentukan path penyimpanan gambar
	fileExtension := filepath.Ext(file.Filename)
	imagePath := fmt.Sprintf("%s/%s%s", imageFolder, slug.Make(input.Title), fileExtension)

	// Simpan gambar ke server
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Membuat slug untuk post
	generatedSlug := slug.Make(input.Title)

	// Simpan data post ke database, termasuk path gambar
	post := models.Post{
		Title:       input.Title,
		Slug:        generatedSlug,
		CategoryID:  input.CategoryID,
		UserID:      input.UserID,
		Description: input.Description,
		Content:     input.Content,
		Image:       imagePath,
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
