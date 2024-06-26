package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// User is a struct
type User struct {
	ID   int
	Name string
}

type Book struct {
	ID              int
	Title           string
	Author          string
	PublicationYear int
	Genre           string
	Status          string
}

type Loan struct {
	ID         int
	BookID     int
	UserID     int
	LoanDate   string
	ReturnDate string
}

var users []User
var books []Book
var loans []Loan

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Library",
		})
	})

	router.GET("/books", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books.html", gin.H{
			"title": "Books",
			"books": books,
		})
	})

	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"title": "Users",
			"users": users,
		})
	})

	router.GET("/loans", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loans.html", gin.H{
			"title": "Loans",
			"loans": loans,
		})
	})

	router.Run(":8080")
}
