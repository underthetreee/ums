package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/underthetreee/ums/internal/model"
)

type UserService interface {
	Register(ctx context.Context, input model.RegisterUserParams) (string, error)
	Login(ctx context.Context, params model.LoginUserParams) (string, error)
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
	var params model.RegisterUserParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	token, err := h.service.Register(r.Context(), params)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var params model.LoginUserParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	token, err := h.service.Login(r.Context(), params)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
