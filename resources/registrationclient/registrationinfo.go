package registrationclient

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/credentials"

	"github.com/wolf1996/gateway/regserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/wolf1996/gateway/token"
	"github.com/golang/protobuf/proto"
	"encoding/base64"
)

type Config struct {
	Addres string
	Crt    string
}

type RegistrationInfo struct {
	Id      int64
	UserId  string
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

func AddRegistration(userId string, eventId int64, token token.Token) (infoV RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
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

func GetRegistrationInfo(id int64, token token.Token) (infoV RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
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

func RemoveRegistration(id int64, token token.Token) (infoV RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	btTok, err := proto.MarshalMessageSetJSON(&token)
	if err != nil {
		log.Print(err.Error())
		return
	}
	strTok := base64.StdEncoding.EncodeToString(btTok)
	md := metadata.Pairs("token", strTok)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
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

func GetRegistrations(id string, pageNum int64, pageSize int64, token token.Token) (info []RegistrationInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := regserver.NewRegistrationServiceClient(conn)
	infoStr, err := cli.GetUserRegistrations(context.Background(), &regserver.UsersRegistrationsRequest{id, pageSize, pageNum})
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
