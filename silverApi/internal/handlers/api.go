package handlers

import (
	"silverApi/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Options("/account/password", AuthorizationOptions)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Options("/new", AuthorizationOptions)
		router.Get("/password", GetPasswordDatabase)
		router.Post("/new", PostUser)
	})
}
