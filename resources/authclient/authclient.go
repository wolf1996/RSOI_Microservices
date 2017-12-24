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