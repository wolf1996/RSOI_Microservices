package application

import (
	"log"
	"gopkg.in/mgo.v2"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"net"
	"github.com/wolf1996/stats/server"
)

type Config struct {
	Port    string
	Crt     string
	Key     string
	MgoConf MongoConfig
}

type MongoConfig struct {
	Addres string
	DbName string
}
var(
	mgoDb *mgo.Database
)
func StartApp(config Config)(err error){
	log.Printf("Mongo database addres userd %v, database is %v", config.MgoConf.Addres, config.MgoConf.DbName)
	session, err := mgo.Dial(config.MgoConf.Addres)
	if err != nil {
		return
	}
	session.Ping()
	if err != nil {
		return
	}
	session.SetSafe(&mgo.Safe{})
	mgoDb = session.DB(config.MgoConf.DbName)
	port := config.Port
	creds, err := credentials.NewServerTLSFromFile(config.Crt, config.Key)
	lis, err := net.Listen("tcp", port)
	log.Printf("Starting on %s", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	server.RegisterStatisticServiceServer(grpcServer, &GrpcServerInstance{})
	grpcServer.Serve(lis)
}