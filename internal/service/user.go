package service

import (
	"context"

	"github.com/underthetreee/ums/internal/domain"
	"github.com/underthetreee/ums/internal/repository"
)

var _ User = &UserService{}

type User interface {
	Register(ctx context.Context, input domain.UserRegisterInput) error
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, input domain.UserRegisterInput) error {
	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
