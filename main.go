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

	// Rute untuk user
	router.GET("/api/users", controllers.Findusers)
	router.POST("/api/users", controllers.StoreUser)
	router.GET("/api/users/:id", controllers.FindUserById)
	router.PUT("/api/users/:id", controllers.UpdateUser)
	router.DELETE("/api/users/:id", controllers.DeleteUser)

	// Rute untuk kategori
	router.GET("/api/categories", controllers.FindCategories)
	router.GET("/api/categories/:id", controllers.FindCategoryById)

	// Rute yang dilindungi dengan middleware AuthRequired
	protected := router.Group("/api/categories")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/", controllers.StoreCategory)
		protected.PUT("/:id", controllers.UpdateCategory)
		protected.DELETE("/:id", controllers.DeleteCategory)
	}

	// Menambahkan middleware pada rute yang membutuhkan autentikasi
	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.AuthRequired())
	{
		// Rute yang memerlukan token autentikasi
		authRoutes.GET("/user", controllers.GetUser)          // Rute untuk mendapatkan user yang sedang login
		authRoutes.POST("/refresh", controllers.RefreshToken) // Rute untuk refresh token
		authRoutes.POST("/logout", controllers.Logout)        // Rute untuk logout
	}

	// Rute untuk login
	router.POST("/api/login", controllers.Login)

	// Jalankan server di port 3000
	router.Run(":3000")
}
