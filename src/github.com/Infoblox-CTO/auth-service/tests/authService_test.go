package authService

import (
	"context"
	"github.com/Infoblox-CTO/auth-service"
	"testing"
)

// Should be used only for unit tests
type InMemoryStorage struct {
	storage map[string]string
}

func (storage InMemoryStorage) IsExist(ctx context.Context, credentials authService.UserCredentials) error {
	password, ok := storage.storage[credentials.UserName]
	if !ok {
		return authService.ErrorNotFound
	}

	if password != credentials.Password {
		return authService.ErrorUnauthorized
	}
	return nil
}

func TestAuthenticate(t *testing.T) {
	inMemoryStorage := InMemoryStorage{storage: make(map[string]string)}
	service := authService.CreateAuthService(inMemoryStorage)

	ResponseErrorUnauthorized := `{"status": ERROR, "result": [], "status_detail" : {"errors": ["Unauthorized"]}}")`
	ResponseOk := `{"status": OK, "result": [{"token" : AliceAlice pwd, "expirity" : time}]}")`

	credentialsUnderTest := authService.UserCredentials{UserName: "Alice", Password: "Alice pwd"}
	ctx := context.Background()
	{
		// User is not registered

		response, error := service.Authenticate(ctx, credentialsUnderTest)
		if error != authService.ErrorUnauthorized {
			t.Errorf("expected error %s but received %s", authService.ErrorUnauthorized, error)
		}
		if response.Response != ResponseErrorUnauthorized {
			t.Errorf("expected response %s but received %s", ResponseErrorUnauthorized, response.Response)
		}

	}

	{
		// User is registered but pwd is wrong
		inMemoryStorage.storage["Alice"] = "wrong pwd"
		response, error := service.Authenticate(ctx, credentialsUnderTest)
		if error != authService.ErrorUnauthorized {
			t.Errorf("expected error %s but received %s", authService.ErrorUnauthorized, error)
		}

		if response.Response != ResponseErrorUnauthorized {
			t.Errorf("expected response %s but received %s", ResponseErrorUnauthorized, response.Response)
		}

	}

	{
		// User is registered everything is ok
		inMemoryStorage.storage[credentialsUnderTest.UserName] = credentialsUnderTest.Password
		response, error := service.Authenticate(ctx, credentialsUnderTest)
		if error != nil {
			t.Errorf("expected no error but received %s", error)
		}

		if response.Response != ResponseOk {
			t.Errorf("expected response %s but received %s", ResponseOk, response.Response)
		}

	}

}
