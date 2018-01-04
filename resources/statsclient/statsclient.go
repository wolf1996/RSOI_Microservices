package statsclient

import (
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"
	"io"
	"google.golang.org/grpc"
	"github.com/wolf1996/gateway/statserver"
	"github.com/wolf1996/gateway/token"
	"github.com/wolf1996/gateway/resources"
)


type Config struct {
	Addres string
	Crt    string
}

type LoginEvent struct {
	Ok 	 	 bool	   `json:"ok" bson:"ok"`
	Info     string    `json:"info" bson:"info"`
}

type ViewEvent struct {
	Path 	 string	   `json:"path" bson:"path"`
	UserId   int64	   `json:"user_id" bson:"user_id"`
}

type ChangeEvent struct {
	Path 	 string	   `json:"path" bson:"path"`
	UserId   int64	   `json:"user_id" bson:"user_id"`
} 
var addres string
var creds credentials.TransportCredentials
var ConnectionError = fmt.Errorf("Can't connect to Statistics")

func SetConfigs(config Config) {
	addres = config.Addres
	log.Print(fmt.Sprintf("used to statistic service %s", addres))
	var err error
	creds, err = credentials.NewClientTLSFromFile(config.Crt, "")
	log.Printf("Used sertificates: %s", config.Crt)
	if err != nil {
		panic(err.Error())
	}
}

func GetLogins( token token.Token) (info []LoginEvent, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := statserver.NewStatisticServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	infoStr, err := cli.GetLogins(ctx,&statserver.Empty{})
	if err != nil {
		log.Print(err)
		return
	}
	var inf *statserver.LoginEvent
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
		info = append(info, LoginEvent{Ok:inf.Ok, Info:inf.Result})
	}
	return
}


func GetViewEvents( token token.Token) (info []ViewEvent, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := statserver.NewStatisticServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	infoStr, err := cli.GetViews(ctx,&statserver.Empty{})
	if err != nil {
		log.Print(err)
		return
	}
	var inf *statserver.ViewInfo
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
		info = append(info, ViewEvent{Path:inf.Path, UserId:inf.UserLogin})
	}
	return
}

func GetChangeEvents( token token.Token) (info []ChangeEvent, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := statserver.NewStatisticServiceClient(conn)
	ctx, err := resources.TokenToContext(token)
	if err != nil {
		return
	}
	infoStr, err := cli.GetChanges(ctx,&statserver.Empty{})
	if err != nil {
		log.Print(err)
		return
	}
	var inf *statserver.ChangeInfo
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
		info = append(info, ChangeEvent{Path:inf.Path, UserId:inf.UserLogin})
	}
	return
}


