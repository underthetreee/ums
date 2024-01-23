package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/underthetreee/ums/internal/model"
)

type UserService interface {
	Register(ctx context.Context, input model.UserRegisterInput) error
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input model.UserRegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
	}
	if err := h.service.Register(context.Background(), input); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
