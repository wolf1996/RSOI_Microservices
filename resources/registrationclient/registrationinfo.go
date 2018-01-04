package registrationclient

import (
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/credentials"

	"github.com/wolf1996/gateway/regserver"
	"google.golang.org/grpc"
	"github.com/wolf1996/gateway/token"
	"github.com/wolf1996/gateway/resources"
)

type Config struct {
	Addres string
	Crt    string
}

type RegistrationInfo struct {
	Id      int64
	UserId  int64
	EventId int64
}

var addres string
var creds credentials.TransportCredentials
var ConnectionError = fmt.Errorf("Can't connect to Registrations")

func SetConfigs(config Config) {
	addres = config.Addres
	log.Print(fmt.Sprintf("used to reg service %s", addres))
	var err error
	creds, err = credentials.NewClientTLSFromFile(config.Crt, "")
	if err != nil {
		panic(err.Error())
	}
}

func AddRegistration(userId int64, eventId int64, token token.Token) (infoV RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	info, err := cli.AddRegistration(ctx, &regserver.RegistrationToAdd{userId, eventId})
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

func GetRegistrationInfo(id int64, token token.Token) (infoV RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	info, err := cli.GetRegistrationInfo(ctx, &regserver.RegistrationId{id})
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

func RemoveRegistration(id int64, token token.Token) (infoV RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
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

func GetRegistrations(id int64, pageNum int64, pageSize int64, token token.Token) (info []RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	infoStr, err := cli.GetUserRegistrations(ctx, &regserver.UsersRegistrationsRequest{id, pageSize, pageNum})
	if err != nil {
		log.Print(err)
		return
	}
	var inf *regserver.RegistrationInfo
	for {
		inf, err = infoStr.Recv()
		if err != nil {
			if err != io.EOF {
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
