package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"

	"github.com/Infoblox-CTO/auth-service"
	proto "github.com/Infoblox-CTO/auth-service/proto"
	authstorage "github.com/Infoblox-CTO/auth-service/storage"
	"github.com/go-kit/kit/endpoint"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	listen, error := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if error != nil {
		grpclog.Fatalf("failed to listen: %v", error)
	}

	inMemoryStorage := authstorage.CreateStorage()
	inMemoryStorage.Register(authService.UserCredentials{UserName: "Alice", Password: "pwd"})
	service := authService.CreateAuthService(inMemoryStorage)

	var authEndpoint endpoint.Endpoint
	authEndpoint = authService.MakeAuthEndpoint(service)
	endpoints := authService.Endpoints{AuthenticateEndpoint: authEndpoint}

	grpcServer := grpc.NewServer()
	serverGRPC := authService.CreateGRPCServer(endpoints)
	proto.RegisterAuthServer(grpcServer, serverGRPC)
	grpcServer.Serve(listen)
}
