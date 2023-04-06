package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/", app.Home)
	mux.Get("/players", app.GetPlayers)
	mux.Get("/players/{id}", app.OnePlayer)

	return mux
}
