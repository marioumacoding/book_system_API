package models

import "gorm.io/gorm"

// Book struct represents the books table in the database
type Book struct {
	gorm.Model
	Title    string
	Author   string
	Quantity int
}
