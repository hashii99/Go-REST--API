package book

import (
	"test/database"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:'Author'`
	Author string `json:'author'`
	Rating int `json: 'rating'`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(&books)
	// c.Send("All Books")

}

func GetSingleBook(c *fiber.Ctx) {
	// c.Send("Get a single book")
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	// c.Send("Add a new book")
	db := database.DBConn
	// var book Book
	// book.Title = "1984"
	// book.Author = "George Orwell"
	// book.Rating = 5 //hard coded

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	// c.Send("Delete a book")
	id := c.Params("id")
	db := database.DBConn

	var book Book 
	db.First(&book, id)
	if book.Title == "" {
		c.Status(503).Send("No book found with given ID")
		return
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")
}
