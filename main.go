package main

import (
	"be-golang/connection"
	"be-golang/controllers"
	"be-golang/middleware" // Impor middleware AuthRequired
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {
	router := gin.Default()

	// Koneksi ke database
	connection.ConnectDatabase()

	// Menyajikan file statis (gambar) dari folder src/images
	router.Static("/src/images", "./src/images")

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
		authRoutes.POST("/post", controllers.StorePost)
		authRoutes.POST("/menus", controllers.StoreMenu)
		authRoutes.PUT("/menus/:id", controllers.UpdateMenu)
		authRoutes.DELETE("/menus/:id", controllers.DeleteMenu)
		authRoutes.POST("/dataweb", controllers.StoreDataWeb)
		authRoutes.PUT("/dataweb/:id", controllers.UpdateDataWeb)
		authRoutes.DELETE("/dataweb/:id", controllers.DeleteDataWeb)
		authRoutes.POST("/sosmed", controllers.StoreSosmed)
		authRoutes.PUT("/sosmed/:id", controllers.UpdateSosmed)
		authRoutes.DELETE("/sosmed/:id", controllers.DeleteSosmed)
	}

	router.POST("/api/login", controllers.Login)
	router.GET("/api/categories", controllers.FindCategories)
	router.GET("/api/categories/:id", controllers.FindCategoryById)
	router.GET("/api/post", controllers.FindPost)
	router.GET(("/api/menus"), controllers.FindMenu)
	router.GET("/api/menus/:id", controllers.FindMenuById)
	router.GET("/api/dataweb", controllers.FindDataWeb)
	router.GET("/api/dataweb/:id", controllers.FindDataWebById)
	router.GET("/api/sosmed", controllers.FindSosmed)
	router.GET("/api/sosmed/:id", controllers.FindSosmedById)

	// Jalankan server di port 3000
	router.Run(":3000")
}
