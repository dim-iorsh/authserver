package grpcclient

import (
	"github.com/Infoblox-CTO/auth-service"
	proto "github.com/Infoblox-CTO/auth-service/proto"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) authService.AuthService {

	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = grpctransport.NewClient(
			conn,
			"authServiceProto.Auth",
			"Authenticate",
			authService.EncodeGRPCAuthenticateRequest,
			authService.DecodeGRPCAuthenticateResponse,
			proto.AuthReply{},
		).Endpoint()
	}

	return authService.Endpoints{AuthenticateEndpoint: authEndpoint}
}
