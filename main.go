package main

import (
	"be-golang/connection"
	"be-golang/controllers"
	"be-golang/middleware" // Impor middleware AuthRequired
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Koneksi ke database
	connection.ConnectDatabase()

	// Rute untuk halaman utama
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang di Blog Paytrizz API")
	})

	// Rute yang memerlukan token autentikasi (untuk pengguna dan kategori)
	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.AuthRequired()) // Menambahkan middleware untuk semua rute dalam grup ini
	{
		// Rute yang memerlukan token autentikasi untuk pengguna
		authRoutes.GET("/users", controllers.Findusers)
		authRoutes.GET("/users/:id", controllers.FindUserById)
		authRoutes.POST("/users", controllers.StoreUser)
		authRoutes.PUT("/users/:id", controllers.UpdateUser)
		authRoutes.DELETE("/users/:id", controllers.DeleteUser)
		authRoutes.GET("/user", controllers.GetUser)
		authRoutes.POST("/refresh", controllers.RefreshToken)
		authRoutes.POST("/logout", controllers.Logout)
		authRoutes.POST("/categories", controllers.StoreCategory)
		authRoutes.PUT("/categories/:id", controllers.UpdateCategory)
		authRoutes.DELETE("/categories/:id", controllers.DeleteCategory)
	}

	router.POST("/api/login", controllers.Login)
	router.GET("/api/categories", controllers.FindCategories)
	router.GET("/api/categories/:id", controllers.FindCategoryById)

	// Jalankan server di port 3000
	router.Run(":3000")
}
