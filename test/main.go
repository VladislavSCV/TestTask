package main

import (
	"fmt"
	"log"

	"github.com/VladislavSCV/TestTask/internal/db"
)

func main() {
	fmt.Print("Hello, world!")
	conn, err := db.ConnToDb()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(conn.Ping())
}
