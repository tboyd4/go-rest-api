package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tboyd4/go-rest-api/models"
)

func FindBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{ "data": books })
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{ "data": book })
}
