package main

import (
	"fmt"
	"test/book"
	"test/database"

	"github.com/gofiber/fiber"
	// "gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// func helloWorld(c *fiber.Ctx) {
// 	c.Send("Hello, Hashini")
// }

func setupRoute(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetSingleBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}
// initiallize db connection
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connection successfully opened!")

	// automigration
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}



func main() {
	
	app := fiber.New() //instantiate a new fiber application
	// app.Get("/") - http get request and mapped to helloWorld named fuction
	// app.Get("/", helloWorld)

	initDatabase()
	defer database.DBConn.Close()

	setupRoute(app)

	app.Listen(3000)
}