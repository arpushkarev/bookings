package main

import (
	"net/http"

	"github.com/arpushkarev/bookings/pkg/config"
	"github.com/arpushkarev/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(cfg *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/about", handlers.Repo.AboutPage)
	mux.Get("/", handlers.Repo.HomePage)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("static", fileServer))

	return mux
}
