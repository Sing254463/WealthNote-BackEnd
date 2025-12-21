package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func ConnectSQLServer() {
    var err error
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
        os.Getenv("SQLSERVER_HOST"),
        os.Getenv("SQLSERVER_USER"),
        os.Getenv("SQLSERVER_PASSWORD"),
        os.Getenv("SQLSERVER_DB"),
    )

    db, err = sql.Open("mssql", connString)
    if err != nil {
        log.Fatalf("Error opening connection to SQL Server: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging SQL Server: %v", err)
    }

    log.Println("Connected to SQL Server!")
}

func GetDB() *sql.DB {
    return db
}