package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)

	mux.Post("/authenticate", app.Authenticate)
	mux.Get("/refresh", app.RefreshToken)
	mux.Get("/logout", app.Logout)

	mux.Get("/movies", app.AllMovies)

	return mux
}
