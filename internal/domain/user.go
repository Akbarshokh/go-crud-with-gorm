package domain

import "context"

type User struct {
	ID       string
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, userID string) (User, error)
}
