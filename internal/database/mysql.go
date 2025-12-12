package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	var err error
	mysqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
	}

	if err = mysqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging MySQL: %v", err)
	}

	log.Println("Connected to MySQL database successfully")
}

func GetMySQLDB() *sql.DB {
	return mysqlDB
}
