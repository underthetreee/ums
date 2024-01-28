package repository

import (
	"context"

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

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
`

func (r *UserRepository) Create(ctx context.Context, user model.User) error {
	// temp
	r.db.MustExec(schema)

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
