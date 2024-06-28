package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/VladislavSCV/TestTask/internal/db"
)

type User struct {
	UserName    string `form:"userName" json:"userName"`
	ContactInfo string `form:"contactInfo" json:"contactInfo"`
}

type Book struct {
	Title           string `form:"title" json:"title"`
	Author          string `form:"author" json:"author"`
	PublicationYear int    `form:"publication-year" json:"publicationYear"`
	Genre           string `form:"genre" json:"genre"`
	Status          string `form:"status" json:"status"`
}

type Loan struct {
	BookID     int    `form:"bookId" json:"bookId"`
	UserID     int    `form:"userId" json:"userId"`
	LoanDate   string `form:"loanDate" json:"loanDate"`
	ReturnDate string `form:"returnDate" json:"returnDate"`
}

var users []User
var books []Book
var loans []Loan

func main() {

	conn, err := db.ConnToDb()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(conn.Ping())

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Library",
		})
	})

	router.GET("/books", func(c *gin.Context) {
		var books []Book
		rows, err := conn.Query("SELECT bookid, title, author, publicationyear, genre, status FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Status)
			if err != nil {
				log.Printf("Error scanning rows: %v", err)
				continue
			}
			books = append(books, book)
		}
		if err = rows.Err(); err != nil {
			log.Fatalf("Error retrieving books: %v", err)
		}
		c.HTML(http.StatusOK, "books.html", gin.H{
			"title": "Books",
			"books": books,
		})
	})

	router.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.BindJSON(&book); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		query := "INSERT INTO books (title, author, publicationyear, genre, status) VALUES(?, ?, ?, ?, ?)"
		_, err := conn.Exec(query, book.Title, book.Author, book.PublicationYear, book.Genre, book.Status)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusOK, "http://127.0.0.1:8000/books")
	})

	router.GET("/users", func(c *gin.Context) {
		var users User[]
		rows, err := conn.Query("SELECT bookid, title, author, publicationyear, genre, status FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Status)
			if err != nil {
				log.Printf("Error scanning rows: %v", err)
				continue
			}
			books = append(books, book)
		}
		if err = rows.Err(); err != nil {
			log.Fatalf("Error retrieving books: %v", err)
		}
		c.HTML(http.StatusOK, "users.html", gin.H{
			"title": "Users",
			"users": users,
		})
	})

	router.GET("/loans", func(c *gin.Context) {
		var loans Loan[]
		rows, err := conn.Query("SELECT bookid, title, author, publicationyear, genre, status FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Status)
			if err != nil {
				log.Printf("Error scanning rows: %v", err)
				continue
			}
			books = append(books, book)
		}
		if err = rows.Err(); err != nil {
			log.Fatalf("Error retrieving books: %v", err)
		}
		c.HTML(http.StatusOK, "loans.html", gin.H{
			"title": "Loans",
			"loans": loans,
		})
	})

	router.Run(":8000")
}
