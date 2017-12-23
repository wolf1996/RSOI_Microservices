package application

import (
	"log"
	"github.com/wolf1996/auth/application/storage"
	_ "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/credentials"
	"github.com/wolf1996/auth/application/models"
	"net"
	"google.golang.org/grpc"
	"github.com/wolf1996/auth/token"
	"github.com/wolf1996/auth/application/tokenanager"
)

type Config struct {
	Port         string
	Crt          string
	Key          string
	StorConfig   storage.Config
	DatabaseConf models.DatabaseConfig
	TokenConf    tokenanager.Config

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

func StartApplication(config Config){
	applyConfig(config)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting on %s", port)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	token.RegisterAuthServiceServer(grpcServer, &AuthServerInstance{})
	grpcServer.Serve(lis)
}