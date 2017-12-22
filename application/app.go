package application

import (
	"log"
	"net"

	"github.com/wolf1996/events/application/models"
	"github.com/wolf1996/events/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Config struct {
	Port         string
	Crt          string
	Key          string
	DatabaseConf models.DatabaseConfig
}

var port string
var creds credentials.TransportCredentials

func applyConfig(config Config) {
	port = config.Port
	models.ApplyConfig(config.DatabaseConf)
	var err error
	creds, err = credentials.NewServerTLSFromFile(config.Crt, config.Key)
	if err != nil {
		panic(err.Error())
	}
}

func StartApplication(config Config) {
	applyConfig(config)
	applyConfig(config)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting on %s", port)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	server.RegisterEventServiceServer(grpcServer, &GrpcEventsServerInstance{})
	grpcServer.Serve(lis)
}
