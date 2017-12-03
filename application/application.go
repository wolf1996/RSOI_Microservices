package application

import (
	"github.com/wolf1996/user/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/wolf1996/user/application/models"
)

type Config struct {
	Port string
	DatabaseConf models.DatabaseConfig
}

var port string

func applyConfig(config Config){
	port = config.Port
	models.ApplyConfig(config.DatabaseConf)
}

func StartApplication(config Config){
	applyConfig(config)
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting on %s", port)
	grpcServer := grpc.NewServer()
	server.RegisterUserServiceServer(grpcServer, &UserInfoServer{})
	grpcServer.Serve(lis)
}
