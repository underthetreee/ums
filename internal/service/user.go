package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/underthetreee/ums/internal/auth"
	"github.com/underthetreee/ums/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, params model.RegisterUserParams) (string, error) {
	existingUser, err := s.repo.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	if existingUser != nil {
		return "", errors.New("user already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := model.User{
		ID:       uuid.New(),
		Name:     params.Name,
		Email:    params.Email,
		Password: string(hashedPass),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID.String(), 1*time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) Login(ctx context.Context, params model.LoginUserParams) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := auth.GenerateToken(user.ID.String(), 1*time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) GetProfile(ctx context.Context) (*model.UserProfileParams, error) {
	id, err := auth.GetUserIDFromClaims(ctx)
	if err != nil {
		return nil, err
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUserByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	userProfile := model.UserProfileParams{
		Name:  user.Name,
		Email: user.Email,
	}
	return &userProfile, nil
}
