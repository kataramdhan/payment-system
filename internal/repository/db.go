package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=password dbname=payment_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")

	return db
}
