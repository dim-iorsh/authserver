package authService

import (
	"context"
	"fmt"
)

type authServiceImpl struct {
	credentialsStorage CredentialsStorage
}

// TODO not for production
const (
	StatusOk             = "OK"
	StatusError          = "ERROR"
	ResponseAuthOkTpl    = `{"status": %s, "result": [{"token" : %s, "expirity" : time}]}")`
	ResponseAuthErrorTpl = `{"status": %s, "result": [], "status_detail" : {"errors": ["%s"]}}")`
	ResponseTokenOkTpl   = `{"status": %s, "result": [{"token" : %s, "expirity" : time}]}")`
)

func (service authServiceImpl) Authenticate(ctx context.Context, credentials UserCredentials) (AuthResponse, error) {
	error := service.credentialsStorage.IsExist(ctx, credentials)
	// TODO: storage errors to authService errors
	if error != nil {
		return AuthResponse{Response: fmt.Sprintf(ResponseAuthErrorTpl, StatusError, ErrorUnauthorized)}, ErrorUnauthorized
	}

	tokenInfo := generateToken(credentials)
	return AuthResponse{Response: fmt.Sprintf(ResponseAuthOkTpl, StatusOk, tokenInfo.token)}, nil
}

// TODO currently it is STUB
func (service authServiceImpl) CheckToken(ctx context.Context, token TokenInfo) (AuthResponse, error) {
	return AuthResponse{Response: fmt.Sprintf(ResponseTokenOkTpl, StatusOk, token.token)}, nil
}

func CreateAuthService(storage CredentialsStorage) AuthService {
	return authServiceImpl{credentialsStorage: storage}
}

func generateToken(credentials UserCredentials) TokenInfo {
	return TokenInfo{token: credentials.UserName + credentials.Password}
}
