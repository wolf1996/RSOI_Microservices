package application

import (
	"github.com/wolf1996/user/server"
	"google.golang.org/grpc"
	"log"
	"net"
)



func StartApplication(){
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.RegisterUserServiceServer(grpcServer, &UserInfoServer{})
	grpcServer.Serve(lis)
}
