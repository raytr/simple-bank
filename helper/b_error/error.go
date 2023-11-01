package b_error

import (
	"errors"
)

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
	ErrUnauthorized = errors.New("unauthorized")
)
