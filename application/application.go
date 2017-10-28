package application

import (
	"github.com/wolf1996/user/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/wolf1996/user/application/models"
)

type UserConfig struct {
	Port string
	DatabaseConf models.UserDatabaseConfig
}

var port string

func applyConfig(config UserConfig){
	port = config.Port
	models.ApplyConfig(config.DatabaseConf)
}

func StartApplication(config UserConfig){
	applyConfig(config)
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.RegisterUserServiceServer(grpcServer, &UserInfoServer{})
	grpcServer.Serve(lis)
}
