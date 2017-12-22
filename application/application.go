package application

import (
	"log"
	"net"

	"github.com/wolf1996/registration/application/models"
	"github.com/wolf1996/registration/server"
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
	lis, err := net.Listen("tcp", port)
	log.Printf("Starting on %s", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	server.RegisterRegistrationServiceServer(grpcServer, &GprsServerInstance{})
	grpcServer.Serve(lis)
}
