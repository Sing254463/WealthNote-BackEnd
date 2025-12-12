package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var postgresDB *sql.DB

func ConnectPostgres() {
	// DSN format สำหรับ lib/pq ควรใช้แบบนี้
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_USER"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_DATABASE"),
		os.Getenv("POSTGRESQL_SSLMODE"),
	)

	var err error
	postgresDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}

	if err = postgresDB.Ping(); err != nil {
		log.Fatalf("Error pinging PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL database successfully")
}

func GetPostgresDB() *sql.DB {
	return postgresDB
}

func ClosePostgres() {
	if postgresDB != nil {
		if err := postgresDB.Close(); err != nil {
			log.Printf("Error closing PostgreSQL connection: %v", err)
		}
	}
}
