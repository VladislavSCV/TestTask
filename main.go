package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/VladislavSCV/TestTask/internal/db"
)

type User struct {
	Userid 		int    `db:"userid" json:"userid"`
	UserName    string `form:"name" json:"name"`
	ContactInfo string `form:"contact-info" json:"contact-info"`
}

type Book struct {
	BookId 			int    `form:"bookId" json:"bookId"`
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
			err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Status)
			if err != nil {
				log.Printf("Error scanning rows: %v", err)
				continue
			}
			books = append(books, book)
		}
		if err = rows.Err(); err != nil {
			log.Fatalf("Error retrieving books: %v", err)
		}
		log.Println(books)
		c.HTML(http.StatusOK, "books.html", gin.H{
			"title": "Books",
			"books": books,
		})
	})

	router.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.Bind(&book); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		query := "INSERT INTO books (title, author, publicationyear, genre, status) VALUES($1, $2, $3, $4, $5)"
		_, err := conn.Exec(query, book.Title, book.Author, book.PublicationYear, book.Genre, book.Status)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusFound, "http://127.0.0.1:8000/books")
	})

	router.GET("/users", func(c *gin.Context) {
		var users []User
		query := "SELECT userid, username, contactinfo FROM users"
		rows, err := conn.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var user User
			err := rows.Scan(&user.Userid, &user.UserName, &user.ContactInfo)
			if err != nil {
				log.Printf("Error scanning rows: %v", err)
				continue
			}
			users = append(users, user)
		}
		if err = rows.Err(); err != nil {
			log.Fatalf("Error retrieving users: %v", err)
		}
		log.Println(users)
		c.HTML(http.StatusOK, "users.html", gin.H{
			"title": "Users",
			"users": users,
		})
	})

	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.Bind(&user); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		query := "INSERT INTO users (username, contactinfo) VALUES($1, $2)"
		_, err := conn.Exec(query, user.UserName, user.ContactInfo)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusFound, "http://127.0.0.1:8000/users")
	})

	router.GET("/loans", func(c *gin.Context) {
		var books []Book
		rows, err := conn.Query("SELECT bookid, title, author, publicationyear, genre, status FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Status)
			if err != nil {
				log.Printf("Error scanning rows: %v", err)
				continue
			}
			books = append(books, book)
		}
		if err = rows.Err(); err != nil {
			log.Fatalf("Error retrieving books: %v", err)
		}
		log.Println(books)
		c.HTML(http.StatusOK, "books.html", gin.H{
			"title": "Books",
			"books": books,
		})
	})

	router.POST("/loans", func(c *gin.Context) {
		var book Book
		if err := c.Bind(&book); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		query := "INSERT INTO books (title, author, publicationyear, genre, status) VALUES($1, $2, $3, $4, $5)"
		_, err := conn.Exec(query, book.Title, book.Author, book.PublicationYear, book.Genre, book.Status)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusFound, "http://127.0.0.1:8000/books")
	})

	router.Run(":8000")
}
