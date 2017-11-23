package userclient

import (
	"github.com/wolf1996/user/server"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
)

type Config struct{
	Addres string
}

type UserInfo struct {
	Name string
	Count int64
	Id 	  int64
}

var addres string

func SetConfigs(config Config){
	addres = config.Addres
	log.Print(fmt.Sprintf("used to userInfo service %s", addres))
}

func IncrementEventsCounter(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewUserServiceClient(conn)
	info, err := cli.IncrementUserCounter(context.Background(), &server.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}

func DecrementEventsCounter(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewUserServiceClient(conn)
	info, err := cli.DecrementUserCounter(context.Background(), &server.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}



func GetUserInfo(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewUserServiceClient(conn)
	info, err := cli.GetUserInfo(context.Background(), &server.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}
