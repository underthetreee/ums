package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/underthetreee/ums/internal/domain"
	"github.com/underthetreee/ums/internal/service"
)

type UserHandler struct {
	service service.User
}

func NewUserHandler(service service.User) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input domain.UserRegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
	}
	if err := h.service.Register(context.Background(), input); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
