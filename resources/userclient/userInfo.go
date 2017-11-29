package userclient

import (
	"github.com/wolf1996/gateway/usserver"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
)

type Config struct{
	Addres string
	QuemanagerConfig QConfig
}

type UserInfo struct {
	Name string
	Count int64
	Id 	  int64
}

var addres string

var ConnectionError = fmt.Errorf("Can't connect to Users")

func SetConfigs(config Config){
	addres = config.Addres
	log.Print(fmt.Sprintf("used to userInfo service %s", addres))
	err := ApplyConfig(config.QuemanagerConfig)
	if err != nil {
		panic(err.Error())
	}
}

func IncrementEventsCounter(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		err = ConnectionError
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	info, err := cli.IncrementUserCounter(context.Background(), &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}

func DecrementEventsCounter(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		err = ConnectionError
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	info, err := cli.DecrementUserCounter(context.Background(), &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}


func DecrementEventsCounterAsync(id string) (err  error){
	UserEventsDecrementCounter(id)
	return nil
}


func GetUserInfo(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		err = ConnectionError
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	info, err := cli.GetUserInfo(context.Background(), &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}
