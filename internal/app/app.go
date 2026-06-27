package app

import (
	"goapi/internal/database"
	"goapi/internal/handlers"
	"goapi/internal/repository"
	"goapi/internal/routes"
	service "goapi/internal/services"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {

	r := chi.NewRouter()

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
