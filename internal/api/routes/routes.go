package routes

import (
	"TextSearchAPI/internal/api/v0.1"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v0.1", func(r chi.Router) {
		r.Use(ApiVersionCtx("v0.1"))
		r.Mount("/search", SearchRouter())
	})

	return r
}

func SearchRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/{text}", func(r chi.Router) {
		r.Get("/", v0_1.HandleSearch)
	})
	return r
}
