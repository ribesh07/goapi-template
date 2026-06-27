package handlers

import (
	"encoding/json"
	"net/http"

	service "goapi/internal/service"
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

	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
