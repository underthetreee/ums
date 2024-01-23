package model

import "github.com/google/uuid"

type UserRegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}
