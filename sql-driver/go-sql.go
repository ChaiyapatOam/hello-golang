package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {
	query := `CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME ,
		PRIMARY KEY (id)
	)`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func Insert(db *sql.DB) {
	var username string
	var password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	createdAt := time.Now()

	result, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Println(id)
}

func query(db *sql.DB) {
	var (
		id      int
		subject string
	)
	queryString := "SELECT id,subject FROM course WHERE id = ?"
	if err := db.QueryRow(queryString, 2).Scan(&id, &subject); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, subject)
}

func main() {
	db, err := sql.Open("mysql", "root:ComCamp34!PassDB@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		fmt.Println("Failed to Connect")
	} else {
		fmt.Println("Connected!")
	}
	// creatingTable(db)
	Insert(db)
	// query(db)
}
