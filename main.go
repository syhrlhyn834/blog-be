package main

import (
	"be-golang/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	connection.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang di Blog Paytrizz API")
	})

	router.Run(":3000")

}
