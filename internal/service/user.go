package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/underthetreee/ums/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, input model.UserRegisterInput) error {
	user := model.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
