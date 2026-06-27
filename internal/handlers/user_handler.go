package handlers

import (
	"encoding/json"
	"net/http"

	service "goapi/internal/services"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {

	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users := h.service.GetUsers()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
