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
	GetProfile(ctx context.Context) (*model.UserProfileParams, error)
	UpdateProfile(ctx context.Context, user model.UserProfileParams) error
	DeleteProfile(ctx context.Context) error
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
		JSONError(w, ErrInvalidInput())
		log.Println(err)
		return
	}

	token, err := h.service.Register(r.Context(), params)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, ErrInternalServer())
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusCreated, token)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var params model.LoginUserParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		JSONError(w, ErrInvalidInput())
		log.Println(err)
		return
	}

	token, err := h.service.Login(r.Context(), params)
	if err != nil {
		JSONError(w, ErrInternalServer())
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusOK, token)
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := h.service.GetProfile(r.Context())
	if err != nil {
		JSONError(w, ErrUnauthorized())
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusOK, profile)
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var params model.UserProfileParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		err := ErrInvalidInput()
		log.Println(err)
		return
	}
	if err := h.service.UpdateProfile(r.Context(), params); err != nil {
		JSONError(w, ErrInternalServer())
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusNoContent, nil)
}

func (h *UserHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteProfile(r.Context()); err != nil {
		JSONError(w, ErrInternalServer())
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusNoContent, nil)
}
