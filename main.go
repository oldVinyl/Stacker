package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DBConnect()
}

func DBConnect() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to MySQL!")
}
