package application

import (
	"net"
	"log"
	"github.com/wolf1996/registration/application/models"
	"google.golang.org/grpc"
	"github.com/wolf1996/registration/server"
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.RegisterUserServiceServer(grpcServer, &GprsServerInstance{})
	grpcServer.Serve(lis)
}
