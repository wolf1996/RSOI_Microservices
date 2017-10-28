package application

import (
	"github.com/wolf1996/events/application/models"
	"github.com/wolf1996/events/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Config struct{
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
	applyConfig(config)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.RegisterEventServiceServer(grpcServer,&GrpcEventsServerInstance{})
	grpcServer.Serve(lis)
}
