package models

import (
	"github.com/jinzhu/gorm"
	config "github.com/nikhilrana/Go-LibraryMgmt/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Quantity    int64  `json:"quantity"`
}

func init() {
	config.InitEnvConfigs()
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) AddNewBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	allBooks := []Book{}
	db.Find(&allBooks)
	return allBooks
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var bookById Book
	db := db.Where("ID=?", id).Find(&bookById)
	return &bookById, db
}

func DeleteBookById(id int64) {
	var deletedBook Book
	db.Where("Id=?", id).Delete(deletedBook)
	// return deletedBook
}
