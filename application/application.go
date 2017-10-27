package application

import (
	"github.com/wolf1996/user/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserConfig struct {
	Port string
}

func StartApplication(config UserConfig){
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.RegisterUserServiceServer(grpcServer, &UserInfoServer{})
	grpcServer.Serve(lis)
}
