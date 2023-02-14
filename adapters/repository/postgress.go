package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {

	connStr := "user=postgres dbname=postgres password=000 sslmode=disable host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Minute * 5)

	fmt.Println("Successfully connected!")

	return db
}
