package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	ID         int       `json:"id"`
	Company    string    `json:"company"`
	Position   string    `json:"position"`
	Date       time.Time `json:"date"`
	ResumePath string    `json:"resume_path"`
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Start SQLite3
	db, err := sql.Open("sqlite3", "./applications.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	router.Run(":8080")
}

func createTable(db *sql.DB) {
	query := `
	  CREATE TABLE IF NOT EXISTS applications (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      company TEXT,
      position TEXT,
      date TEXT,
      resume_path TEXT
    );
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
