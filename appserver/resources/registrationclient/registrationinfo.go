package registrationclient

import (
	"github.com/wolf1996/gateway/regserver"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
	"google.golang.org/grpc/metadata"
)

type Config struct{
	Addres string
}

type RegistrationInfo struct {
	Id      int64
	UserId  int64
	EventId int64
}

var addres string

func SetConfigs(config Config){
	addres = config.Addres
	log.Print(fmt.Sprintf("used to reg service %s", addres))
}

func AddRegistration(userId, eventId int64) (infoV RegistrationInfo, err error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := regserver.NewUserServiceClient(conn)
	info, err := cli.AddRegistration(context.Background(), &regserver.RegistrationToAdd{userId, eventId})
	if err != nil {
		return
	}
	infoV = RegistrationInfo{
		info.Id,
		info.UserId,
		info.EventId,
	}
	return
}

func GetRegistrationInfo(id int64) (infoV RegistrationInfo, err error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := regserver.NewUserServiceClient(conn)
	info, err := cli.GetRegistrationInfo(context.Background(), &regserver.RegistrationId{id})
	if err != nil {
		return
	}
	infoV = RegistrationInfo{
		info.Id,
		info.UserId,
		info.EventId,
	}
	return
}

func RemoveRegistration(id int64, md metadata.MD) (infoV RegistrationInfo, err error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	cli := regserver.NewUserServiceClient(conn)
	info, err := cli.RemoveRegistration(ctx, &regserver.RegistrationId{id})
	if err != nil {
		return
	}
	infoV = RegistrationInfo{
		info.Id,
		info.UserId,
		info.EventId,
	}
	return
}