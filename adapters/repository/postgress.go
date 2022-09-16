package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {

	connStr := "user=postgres dbname=postgres password=000 sslmode=disable host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
