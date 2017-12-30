package application

import (
	"log"
 	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"net"
	"github.com/wolf1996/stats/server"
	"github.com/wolf1996/stats/application/model"
)

type Config struct {
	Port    string
	Crt     string
	Key     string
	MgoConf model.MongoConfig
}

func StartApp(config Config)(err error){
	port := config.Port
	model.ApplyConfig(config.MgoConf)
	creds, err := credentials.NewServerTLSFromFile(config.Crt, config.Key)
	lis, err := net.Listen("tcp", port)
	log.Printf("Starting on %s", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	server.RegisterStatisticServiceServer(grpcServer, &GrpcServerInstance{})
	grpcServer.Serve(lis)
	return
}