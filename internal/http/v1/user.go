package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/underthetreee/ums/internal/model"
)

type UserService interface {
	Register(ctx context.Context, input model.UserRegisterInput) (string, error)
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
		log.Println(err)
	}
	token, err := h.service.Register(r.Context(), input)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}
