package authService

import (
	"context"
	"errors"
)

// UserCredentials represents a user data.
type UserCredentials struct {
	UserName string
	Password string
}

type AuthResponse struct {
	Response string
}

type TokenInfo struct {
	token string
}

// AuthService is a base interface for authentication logic.
type AuthService interface {
	// Checks UserCredentials. If UserCredentials is wrong returns (AuthResponse, error),
	// otherwise - (AuthResponse, nil)
	Authenticate(ctx context.Context, credentials UserCredentials) (AuthResponse, error)
	// Check token. If token is not valid (AuthResponse, error),
	// otherwise - (AuthResponse, nil)
	//CheckToken(ctx context.Context, token TokenInfo) (AuthResponse, error)
}

// AuthStorage is a base interface for a storage where user's credentials should be stored.
type CredentialsStorage interface {
	// Checks if UserCredentials exists or not.
	// If UserCredentials doesn't exist return error, otherwise - nil
	IsExist(ctx context.Context, credentials UserCredentials) error
}

var (
	ErrorUnauthorized = errors.New("Unauthorized")
	ErrorNotFound     = errors.New("Not found")
)
