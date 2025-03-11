package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "Book1", Author: "Mariam", Quantity: 1},
	{ID: 2, Title: "Book2", Author: "Marina", Quantity: 3},
	{ID: 3, Title: "Book3", Author: "Micho", Quantity: 4},
	{ID: 4, Title: "Book4", Author: "Ayman", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if _, isFound := searchBookByID(newBook.ID); isFound {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book with given ID already exists"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func searchBookByID(id int) (int, bool) {
	for index, b := range books {
		if b.ID == id {
			return index, true
		}
	}
	return -1, false
}

func getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert string to int
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	if index, isFound := searchBookByID(id); isFound {
		c.IndentedJSON(http.StatusOK, books[index])
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
}

func updateBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var updatedBook book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	if index, isFound := searchBookByID(id); isFound {
		updatedBook.ID = id
		books[index] = updatedBook
		c.IndentedJSON(http.StatusOK, books[index])
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
}

func patchBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var updatedBook map[string]interface{}
	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	if index, isFound := searchBookByID(id); isFound {
		if title, ok := updatedBook["title"].(string); ok {
			books[index].Title = title
		}
		if author, ok := updatedBook["author"].(string); ok {
			books[index].Author = author
		}
		if quantity, ok := updatedBook["quantity"].(float64); ok {
			books[index].Quantity = int(quantity)
		}
		c.IndentedJSON(http.StatusOK, books[index])
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
}

func deleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	if index, isFound := searchBookByID(id); isFound {
		books = append(books[:index], books[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Book with given ID is deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with given ID not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBookByID)
	router.PATCH("/books/:id", patchBookByID)
	router.DELETE("/books/:id", deleteByID)
	router.Run("localhost:8080")
}
