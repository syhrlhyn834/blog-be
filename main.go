package main

import (
	"be-golang/connection"
	"be-golang/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	connection.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang di Blog Paytrizz API")
	})

	router.GET("/api/users", controllers.Findusers)

	router.POST("/api/users", controllers.StoreUser)

	router.GET("/api/users/:id", controllers.FindUserById)

	router.PUT("/api/users/:id", controllers.UpdateUser)

	router.DELETE("/api/users/:id", controllers.DeleteUser)

	router.GET("/api/categories", controllers.FindCategories)

	router.POST("/api/categories", controllers.StoreCategory)

	router.GET("/api/categories/:id", controllers.FindCategoryById)

	router.PUT("/api/categories/:id", controllers.UpdateCategory)

	router.DELETE("/api/categories/:id", controllers.DeleteCategory)

	router.Run(":3000")

}
