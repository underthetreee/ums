package repository

import (
	"context"
	"database/sql"
	"errors"

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
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err

	}
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx, &user, "SELECT * from users WHERE email = $1", email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
