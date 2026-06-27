package app

import (
	"fmt"
	"goapi/internal/database"
	"goapi/internal/handlers"
	"goapi/internal/repository"
	"goapi/internal/routes"
	"goapi/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Go API Server Running")
	})

	// Database
	db := database.Connect()

	// Users
	// Repository
	userRepo := repository.NewUserRepository(db)

	// Service
	userService := service.NewUserService(userRepo)

	// Handler
	userHandler := handlers.NewUserHandler(userService)

	// Routes
	routes.RegisterUserRoutes(r, userHandler)

	return r
}
