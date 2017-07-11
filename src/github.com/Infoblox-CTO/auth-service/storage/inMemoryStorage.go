package authService

import (
	"context"
	auth "github.com/Infoblox-CTO/auth-service"
)

type InMemoryStorage struct {
	storage map[string]string
}

func (storage InMemoryStorage) IsExist(_ context.Context, credentials auth.UserCredentials) error {
	password, ok := storage.storage[credentials.UserName]
	if !ok {
		return auth.ErrorNotFound
	}

	if password != credentials.Password {
		return auth.ErrorUnauthorized
	}
	return nil
}

func (storage InMemoryStorage) Register(credentials auth.UserCredentials) {
	storage.storage[credentials.UserName] = credentials.Password
}

func CreateStorage() InMemoryStorage {
	return InMemoryStorage{storage: make(map[string]string)}
}
