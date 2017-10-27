package resources

import (
	"github.com/wolf1996/gateway/server"
	"google.golang.org/grpc"
	"context"
)

type UserInfoConfig struct{
	UserInfoServiceAddres string
}

type UserInfo struct {
	Name string
	Count int64
}

var addres string

func SetUserInfoConfigs(config UserInfoConfig){
	addres = config.UserInfoServiceAddres
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
	return &UserInfo{info.Name, info.EventsSubscribed}, nil
}
