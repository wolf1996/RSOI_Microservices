package registrationclient

import (
	"github.com/wolf1996/gateway/regserver"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
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