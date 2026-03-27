package client

import "errors"

var (
	ErrAuthKeyRequired      = errors.New("auth key is required")
	ErrScopeRequired        = errors.New("scope is required")
	ErrTokenManagerRequired = errors.New("token manager is required")
	ErrInvalidAuthHeader    = errors.New("invalid authorization header")
)
