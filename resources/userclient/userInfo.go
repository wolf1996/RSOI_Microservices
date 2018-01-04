package userclient

import (
	"fmt"
	"log"

	"github.com/wolf1996/gateway/usserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/wolf1996/gateway/token"
	"github.com/wolf1996/gateway/resources"
)

type Config struct {
	Addres           string
	Crt              string
	QuemanagerConfig QConfig
}

type UserInfo struct {
	Name  string
	Count int64
	Id    int64
}

var addres string
var creds credentials.TransportCredentials
var ConnectionError = fmt.Errorf("Can't connect to Users")

func SetConfigs(config Config) {
	addres = config.Addres
	log.Print(fmt.Sprintf("used to userInfo service %s", addres))
	err := ApplyConfig(config.QuemanagerConfig)
	if err != nil {
		panic(err.Error())
	}
	creds, err = credentials.NewClientTLSFromFile(config.Crt, "")
	if err != nil {
		panic(err.Error())
	}
}

func IncrementEventsCounter(id int64, token token.Token) (uinf *UserInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	info, err := cli.IncrementUserCounter(ctx, &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}

func DecrementEventsCounter(id int64, token token.Token) (uinf *UserInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	info, err := cli.DecrementUserCounter(ctx, &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}

func DecrementEventsCounterAsync(id int64, token token.Token) (err error) {
	UserEventsDecrementCounter(id, token)
	return nil
}

func GetUserInfo(id int64, token token.Token) (uinf *UserInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := usserver.NewUserServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	info, err := cli.GetUserInfo(ctx, &usserver.UserId{id})
	if err != nil {
		return
	}
	return &UserInfo{info.Name, info.EventsSubscribed, info.Id}, nil
}