package main

import (
	"flag"
	"fmt"
	"log"
	"movies-be/internal/repository"
	"movies-be/internal/types"
	"time"

	"net/http"
)

const port = 8080

type application struct {
	Domain       string
	DSN          string
	DB           types.DatabaseRepo
	Auth         auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "asdasdasdasdasdasd", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()

	conn, err := app.connectToDB()

	if err != nil {
		log.Fatal(err)
	}
	app.DB = &repository.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Auth = auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 20,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieDomain:  app.CookieDomain,
		// CookieName:    "__Host-refersh_token",
		CookieName: "refersh_token",
	}

	log.Println("Starting application server on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
