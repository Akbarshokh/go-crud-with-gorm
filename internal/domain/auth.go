package domain

import (
	"context"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, req User) (string, error)
	Login(ctx context.Context, req User) (JwtTokens, error)
}
