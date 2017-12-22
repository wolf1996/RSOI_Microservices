package application

import (
	"github.com/wolf1996/user/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"github.com/wolf1996/user/application/models"
)

type Config struct {
	Port string
	Crt  string
	Key  string
	DatabaseConf models.DatabaseConfig
}

var port string
var creds credentials.TransportCredentials

func applyConfig(config Config){
	port = config.Port
	models.ApplyConfig(config.DatabaseConf)
	var err error
	creds,err = credentials.NewServerTLSFromFile(config.Crt, config.Key)
	if err != nil {
		panic(err.Error())
	}
}

func StartApplication(config Config){
	applyConfig(config)
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting on %s", port)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	server.RegisterUserServiceServer(grpcServer, &UserInfoServer{})
	grpcServer.Serve(lis)
}
