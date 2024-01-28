package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/underthetreee/ums/internal/auth"
	"github.com/underthetreee/ums/internal/model"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) Register(ctx context.Context, input model.UserRegisterInput) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := model.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPass),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := auth.NewToken(user.ID.String(), 1*time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}
