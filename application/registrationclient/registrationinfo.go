package registrationclient

import (
	"github.com/wolf1996/registration/server"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
	"google.golang.org/grpc/metadata"
	"io"
)

type Config struct{
	Addres string
}

type RegistrationInfo struct {
	Id      int64
	UserId  string
	EventId int64
}

var addres string

func SetConfigs(config Config){
	addres = config.Addres
	log.Print(fmt.Sprintf("used to reg service %s", addres))
}

func AddRegistration(userId string,  eventId int64) (infoV RegistrationInfo, err error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewRegistrationServiceClient(conn)
	info, err := cli.AddRegistration(context.Background(), &server.RegistrationToAdd{userId, eventId})
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
	cli := server.NewRegistrationServiceClient(conn)
	info, err := cli.GetRegistrationInfo(context.Background(), &server.RegistrationId{id})
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
	cli := server.NewRegistrationServiceClient(conn)
	info, err := cli.RemoveRegistration(ctx, &server.RegistrationId{id})
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

func GetRegistrations(id string, pageNum int64, pageSize int64)(info []RegistrationInfo, err error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewRegistrationServiceClient(conn)
	infoStr, err := cli.GetUserRegistrations(context.Background(), &server.UsersRegistrationsRequest{id,pageSize,pageNum})
	if err != nil {
		log.Print(err)
		return
	}
	var inf *server.RegistrationInfo
	for {
		inf, err  = infoStr.Recv()
		if err != nil {
			if err != io.EOF{
				log.Print(err.Error())
				return
			}
			err = nil
			infoStr.CloseSend()
			return
		}
		info = append(info, RegistrationInfo{inf.Id, inf.UserId, inf.EventId})
	}
	return
}