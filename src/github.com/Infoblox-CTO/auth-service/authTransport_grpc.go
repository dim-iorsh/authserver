package authService

import (
	"context"
	proto "github.com/Infoblox-CTO/auth-service/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"
)

type grpcServer struct {
	authenticate grpctransport.Handler
}

func CreateGRPCServer(endpoints Endpoints) proto.AuthServer {
	server := grpctransport.NewServer(endpoints.AuthenticateEndpoint, DecodeGRPCAuthenticateRequest, EncodeGRPCAuthenticateResponse)
	return &grpcServer{authenticate: server}
}

func (server *grpcServer) Authenticate(ctx oldcontext.Context, request *proto.AuthRequest) (*proto.AuthReply, error) {
	_, reply, _ := server.authenticate.ServeGRPC(ctx, request)
	return reply.(*proto.AuthReply), nil
}

// DecodeGRPCAuthenticateRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Authenticate request to a user-domain Authenticate request. Primarily useful in a server.
func DecodeGRPCAuthenticateRequest(_ context.Context, grpcRequest interface{}) (interface{}, error) {
	request := grpcRequest.(*proto.AuthRequest)
	return UserCredentials{UserName: request.UserName, Password: request.Password}, nil
}

// EncodeGRPCAuthenticateResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain Authenticate response to a gRPC Authenticate reply. Primarily useful in a server.
func EncodeGRPCAuthenticateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AuthResponse)
	return &proto.AuthReply{Response: resp.Response}, nil
}

// DecodeGRPCAuthenticateResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Authenticate reply to a user-domain Autenticate response. Primarily useful in a client.
func DecodeGRPCAuthenticateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*proto.AuthReply)
	return reply, nil
}

// EncodeGRPCAuthenticateRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain Authenticate request to a gRPC Authenticate request. Primarily useful in a client.
func EncodeGRPCAuthenticateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(UserCredentials)
	return &proto.AuthRequest{UserName: req.UserName, Password: req.Password}, nil
}
