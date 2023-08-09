package models

import (
	"github.com/chandrasekhar-developer/go-bookstore-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model

	Name        string `grom:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db := config.GetDB()
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	db := config.GetDB()
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	db := config.GetDB()
	var book Book
	DB := db.Where("ID=?", Id).Find(&book)
	return &book, DB
}

func DeleteBook(Id int64) Book {
	db := config.GetDB()
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
