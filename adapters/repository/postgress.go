package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	postgresHost := os.Getenv("POSTGRES_HOST")
	connStr := fmt.Sprintf("host=%s port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", postgresHost)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting database: %q", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		return nil
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Minute * 5)

	fmt.Println("Database Successfully connected!")

	return db
}
