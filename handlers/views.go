package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"canvas/views"
)

func Home(mux chi.Router) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_ = views.Home().Render(w)
	})
}

func Team(mux chi.Router) {
	mux.Get("/team", func(w http.ResponseWriter, r *http.Request) {
		_ = views.Team().Render(w)
	})
}

func Projects(mux chi.Router) {
	mux.Get("/projects", func(w http.ResponseWriter, r *http.Request) {
		_ = views.Projects().Render(w)
	})
}

func Calendar(mux chi.Router) {
	mux.Get("/calendar", func(w http.ResponseWriter, r *http.Request) {
		_ = views.Calendar().Render(w)
	})
}
