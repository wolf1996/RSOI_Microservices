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
}

type UserInfo struct {
	Name string
	Count int64
}

var addres string

func SetConfigs(config Config){
	addres = config.Addres
	log.Print(fmt.Sprintf("used to userInfo service %s", addres))
}


func GetUserInfo(id string) (uinf *UserInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	info, err := cli.GetUserInfo(context.Background(), &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed}, nil
}
