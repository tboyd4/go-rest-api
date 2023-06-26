package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tboyd4/go-rest-api/models"
	"github.com/tboyd4/go-rest-api/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "books api"})
	})

	r.GET("/api/books", controllers.FindBooks)
	r.POST("/api/books", controllers.FindBooks)

	models.ConnectDatabase()

	r.Run()
}
