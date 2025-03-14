package databases

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"book_system/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Change username, password, and database name accordingly
	dsn := "root:password@tcp(127.0.0.1:3306)/books_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to the database successfully!")
		// Automatically migrate the Book model (creates 'books' table)
		DB.AutoMigrate(&models.Book{})
}
