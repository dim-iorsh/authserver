package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Infoblox-CTO/auth-service"
	grpcclient "github.com/Infoblox-CTO/auth-service/client"
	"google.golang.org/grpc"
)

func main() {

	var grpcAddr = flag.String("grpc.addr", "localhost:10000", "gRPC (HTTP) address of auth server")
	flag.Parse()

	var (
		service authService.AuthService
		err     error
	)
	if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		defer conn.Close()
		service = grpcclient.New(conn)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	responce, err := service.Authenticate(context.Background(), authService.UserCredentials{UserName: "Alice", Password: "pwd"})

	fmt.Println(responce)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
