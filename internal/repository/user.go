package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/underthetreee/ums/internal/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user model.User) error {
	fmt.Printf("%+v\n", user)
	return nil
}
