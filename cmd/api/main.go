package main

import (
	"flag"
	"fmt"
	"log"
	"movies-be/internal/repository"
	"movies-be/internal/types"

	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     types.DatabaseRepo
}

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()

	if err != nil {
		log.Fatal(err)
	}
	app.DB = &repository.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("Starting application server on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
