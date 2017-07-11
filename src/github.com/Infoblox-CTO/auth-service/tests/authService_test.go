package authService

import (
	"context"
	"github.com/Infoblox-CTO/auth-service"
	authstorage "github.com/Infoblox-CTO/auth-service/storage"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	inMemoryStorage := authstorage.CreateStorage()
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
		inMemoryStorage.Register(authService.UserCredentials{UserName: "Alice", Password: "wrong pwd"})

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
		inMemoryStorage.Register(credentialsUnderTest)
		response, error := service.Authenticate(ctx, credentialsUnderTest)
		if error != nil {
			t.Errorf("expected no error but received %s", error)
		}

		if response.Response != ResponseOk {
			t.Errorf("expected response %s but received %s", ResponseOk, response.Response)
		}

	}

}
