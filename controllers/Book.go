package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/narasimha-1511/zolo-backend/config"
	"github.com/narasimha-1511/zolo-backend/models"
)

func GetBook(c *gin.Context){
    var book []models.Book

    config.DB.Find(&book)
    

    c.JSON(200, gin.H{
        "Books": book,
    })
}

func CreateBookParams(c *gin.Context){
    var book models.Book

    book.BookID = uint64(rand.Int63())
    book.Name = c.Param("name")
    book.Title = c.Param("title")
    book.Author = c.Param("author")

	config.DB.AutoMigrate(&models.Book{})

    config.DB.Create(&book)

    c.JSON(200, gin.H{
        "book_id": book.BookID,
		"name": book.Name,
		"title": book.Title,
		"author": book.Author,
		"status": "Book Created Successfully",
    })
}

func CreateBookPostForm(c *gin.Context){
	var book models.Book

	book.BookID = uint64(rand.Int63())
	book.Name = c.PostForm("name")
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")

	config.DB.Create(&book)

	c.JSON(200, gin.H{
		"book_id": book.BookID,
		"name": book.Name,
		"title": book.Title,
		"author": book.Author,
		"status": "Book Created Successfully",
	})
}

