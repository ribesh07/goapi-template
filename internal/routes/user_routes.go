package routes

import (
	"goapi/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, h *handlers.UserHandler) {

	r.Route("/users", func(r chi.Router) {

		r.Get("/", h.GetUsers)

	})

}
