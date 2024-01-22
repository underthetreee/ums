package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/underthetreee/ums/internal/domain"
)

var _ User = &UserRepo{}

type User interface {
	Create(ctx context.Context, user domain.User) error
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user domain.User) error {
	user.ID = uuid.New()
	fmt.Printf("%+v\n", user)
	return nil
}
