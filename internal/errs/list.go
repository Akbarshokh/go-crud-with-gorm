package errs

import "errors"

var (
	ErrInternal          = errors.New("internal error")
	ErrValidation        = errors.New("validation error")
	ErrExternal          = errors.New("external service is died")
	ErrAuthorization     = errors.New("authorization failed")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrTokenEmpty        = errors.New("empty token")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrEmailAlreadyExist = errors.New("email already exists")
	ErrNotFound          = errors.New("entity not found")
)
