package authclient

import (
	"github.com/wolf1996/gateway/token"
	"log"
	"fmt"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"context"
)

var ConnectionError = fmt.Errorf("Can't connect to Authclient")

type Config struct {
	Addres           string
	Crt              string
}

type LogInData struct {
	LogIn 			 string
	Pass  			 string
}

type ClientInfo struct {
	Id int64
	Name string
	RedirUrl string
}

var (
	addres string
	creds  credentials.TransportCredentials
)

func SetConfigs(config Config) {
	addres = config.Addres
	log.Print(fmt.Sprintf("used to eventInfo service %s", addres))
	var err error
	creds, err = credentials.NewClientTLSFromFile(config.Crt, "")
	if err != nil {
		panic(err.Error())
	}
}


func ValidateAccessToken(instoken string)(tkn token.Token, err error){
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := token.NewAuthServiceClient(conn)
	validationRes, err := cli.AccessTokenValidation(context.Background(), &token.AccessTokenMsg{instoken})
	if err != nil {
		return
	}
	if validationRes.Valid{
		tkn = *validationRes.Tok
	} else {
		err = fmt.Errorf("invalid token")
	}
	return
}

func GetTokenPair(data LogInData)(access string,refresh string,err error){
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := token.NewAuthServiceClient(conn)
	pairetoken, err := cli.GetTokenpair(context.Background(),&token.SignInPair{data.LogIn, data.Pass})
	if err != nil {
		return
	}
	access = pairetoken.AccessToken.TokenString
	refresh = pairetoken.RefreshToken.TokenString
	return
}

func UpdateTokens(reftoken string)(access string,refresh string,err error){
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := token.NewAuthServiceClient(conn)
	pairetoken, err := cli.GetAccessToken(context.Background(), &token.RefreshTokenMsg{reftoken})
	if err != nil {
		return
	}
	access = pairetoken.AccessToken.TokenString
	refresh = pairetoken.RefreshToken.TokenString
	return
}

func GetClientInfo(id int64) (inf ClientInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := token.NewAuthServiceClient(conn)
	clientinfo, err := cli.GetClientInfo(context.Background(),&token.ClientId{id})
	if err != nil {
		return
	}
	inf.Id = clientinfo.Id
	inf.Name = clientinfo.Name
	inf.RedirUrl = clientinfo.Redirurl
	return
}

func GetCodeGrant(tkn token.Token)(code string, err error){
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := token.NewAuthServiceClient(conn)
	cf, err := cli.GetCodeFlow(context.Background(),&tkn)
	if err != nil {
		return
	}
	code = cf.TokenString
	return
}

func GetCodeflowTokenPair(cflow string)(access string,refresh string,err error){
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := token.NewAuthServiceClient(conn)
	pairetoken, err := cli.ShiftCodeFlow(context.Background(),&token.CodeflowTokenMsg{cflow})
	if err != nil {
		return
	}
	access = pairetoken.AccessToken.TokenString
	refresh = pairetoken.RefreshToken.TokenString
	return
}