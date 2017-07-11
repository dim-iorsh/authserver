package authService

import (
	"context"
	proto "github.com/Infoblox-CTO/auth-service/proto"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AuthenticateEndpoint endpoint.Endpoint
}

// Authenticate implements Service. Primarily useful in a client.
func (e Endpoints) Authenticate(ctx context.Context, credentials UserCredentials) (AuthResponse, error) {
	response, _ := e.AuthenticateEndpoint(ctx, credentials)
	return AuthResponse{Response: response.(*proto.AuthReply).Response}, nil
}

// MakeAuthenticateEndpoint returns an endpoint that invokes Authenticate on the service.
// Primarily useful in a server.
func MakeAuthEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		credentials := request.(UserCredentials)
		reply, _ := s.Authenticate(ctx, credentials)
		return reply, nil
	}
}
