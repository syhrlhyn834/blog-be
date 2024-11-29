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

	router.Run(":3000")

}
