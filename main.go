package main

import (
	"book_system/databases"
	"book_system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type book struct {
// 	ID       int    `json:"id"`
// 	Title    string `json:"title"`
// 	Author   string `json:"author"`
// 	Quantity int    `json:"quantity"`
// }

// var books = []book{
// 	{ID: 1, Title: "Book1", Author: "Mariam", Quantity: 1},
// 	{ID: 2, Title: "Book2", Author: "Marina", Quantity: 3},
// 	{ID: 3, Title: "Book3", Author: "Micho", Quantity: 4},
// 	{ID: 4, Title: "Book4", Author: "Ayman", Quantity: 6},
// }

func getBooks(c *gin.Context) {
	var books[]models.Book
	databases.DB.Find(&books)
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := databases.DB.First(&newBook, "id = ?", newBook.ID).Error; err != nil{
		databases.DB.Create(&newBook)
		c.IndentedJSON(http.StatusCreated, newBook)
		return
	}
	
	c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book with given ID already exists"})
}

func getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var book models.Book

	if err := databases.DB.First(&book, "id = ?", id).Error; err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func updateBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert ID from request to int
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var book models.Book

	if err := databases.DB.First(&book, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
		return
	}

	var updates map[string]interface{}

	if err := c.BindJSON(&updates) ; err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}
	
	if err := databases.DB.Model(&book).Updates(updates).Error; err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update book"})
		return
	}
	// we need to refetch the book after updating it to get the correct data.
	databases.DB.First(&book, "id = ?", id)
	c.IndentedJSON(http.StatusOK, book)
}

func deleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}
	var book models.Book

	if err := databases.DB.First(&book, "id = ?", id).Error ; err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
		return
	}
	databases.DB.Delete(&book, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book with given ID is deleted"})
}

func main() {
	databases.ConnectDatabase()
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PATCH("/books/:id", updateBookByID)
	router.PUT("/books/:id", updateBookByID)
	// router.PATCH("/books/:id", patchBookByID)
	router.DELETE("/books/:id", deleteByID)
	router.Run("localhost:8080")
}
