package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func main() {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	a := App{}
	a.Initialize(cache, db)
	a.Run(":8080")
}
