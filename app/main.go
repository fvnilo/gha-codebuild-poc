package main

import (
	"database/sql"
	"fmt"
	"log"
  "os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	// Define MySQL connection string
	connectionString := fmt.Sprintf("testuser:testpassword@tcp(%s:3306)/testdb", host)

	// Open connection to MySQL
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Ping the database to check connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL!")
}
