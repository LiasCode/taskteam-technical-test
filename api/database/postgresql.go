package database

import (
	"database/sql" // add this
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresqlConnection() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("Connected to postgresql database \n")

	return db, nil
}
