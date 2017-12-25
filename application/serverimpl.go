package application

import (
	"github.com/wolf1996/auth/token"
	"github.com/wolf1996/auth/application/tokenanager"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/wolf1996/auth/application/models"
	"log"
	"github.com/wolf1996/auth/application/storage"
	"time"
)

type AuthServerInstance struct {
}
/*
	GetAccessToken(context.Context, *RefreshTokenMsg) (*AccessTokenMsg, error)
	GetTokenpair(context.Context, *SignInPair) (*Tokenpair, error)
	RefreshTokenValidation(context.Context, *RefreshTokenMsg) (*ValidResult, error)
	AccessTokenValidation(context.Context, *AccessTokenMsg) (*ValidResult, error)
 */
func (inst *AuthServerInstance)GetAccessToken(cnt context.Context, rfrsh *token.RefreshTokenMsg) (tkns *token.Tokenpair,err error){
	tkn, err := tokenanager.ValidateRefreshToken(rfrsh.TokenString)
	tkns = &token.Tokenpair{&token.AccessTokenMsg{}, &token.RefreshTokenMsg{}}
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	cntns, err := storage.CheckTokenStorage(rfrsh.TokenString)
	if err != nil{
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	if !cntns{
		log.Printf("Token not in list:%s", rfrsh.TokenString)
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	actoken, _ ,err := tokenanager.RefreshAccessToken(tkn)
	if err != nil {
		log.Printf("Can't refresh access:%s", err.Error())
		err = status.Errorf(codes.Internal, "Token validation failed")
		return
	}
	restoken, accexp,err := tokenanager.RefreshRefreshToken(tkn)
	if err != nil {
		log.Printf("Can't refresh access:%s", err.Error())
		err = status.Errorf(codes.Internal, "Token validation failed")
		return
	}
	storage.RemoveToken(rfrsh.TokenString)
	storage.AddTokenStorage(restoken, time.Unix(accexp,0))
	tkns.RefreshToken = &token.RefreshTokenMsg{ restoken}
	tkns.AccessToken  = &token.AccessTokenMsg{ actoken}
	return
}

func (inst *AuthServerInstance)GetTokenpair(cnt context.Context,spar *token.SignInPair) (tkns *token.Tokenpair,err error){
	tkns = &token.Tokenpair{&token.AccessTokenMsg{}, &token.RefreshTokenMsg{}}
	uinf, err := models.CheckPass(models.LogIn{Login:spar.Login,Pass:spar.Pass})
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		if err == models.NotFound {
			err = status.Errorf(codes.NotFound, "Invalid Login Password")
		} else {
			err = status.Errorf(codes.Internal, "Server Error")
		}
		return
	}
	tkns.AccessToken.TokenString, _, err = tokenanager.ProduceAccessToken(uinf)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	var rf int64
	tkns.RefreshToken.TokenString, rf, err = tokenanager.ProduceRefreshToken(uinf)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	storage.AddTokenStorage(tkns.RefreshToken.TokenString, time.Unix(rf,0))
	return
}

func (inst *AuthServerInstance)AccessTokenValidation(cnt context.Context,at *token.AccessTokenMsg) (val *token.ValidResult,err error){
	val = &token.ValidResult{Valid:false, Tok:&token.Token{}}
	valu, err := tokenanager.ValidateAccessToken(at.TokenString)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Tiken validation failed")
	} else {
		val.Valid = true
	}
	val.Tok = &token.Token{valu.UserId, valu.LogIn}
	return
}

//GetCodeFlow(context.Context, *RefreshTokenMsg) (*CodeflowTokenMsg, error)
//ShiftCodeFlow(context.Context, *CodeflowTokenMsg) (*Tokenpair, error)

func (inst *AuthServerInstance) GetCodeFlow(cnt context.Context,rt *token.Token) (cf *token.CodeflowTokenMsg, err error) {
	cf = &token.CodeflowTokenMsg{}
	cfls,exp,err := tokenanager.NewCodeFlow(rt.Id, rt.LogIn)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	err = storage.AddTokenStorage(cfls, time.Unix(exp,0))
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.Internal, "Redis storage error")
		return
	}
	cf.TokenString = cfls
	return
}

func  (inst *AuthServerInstance)ShiftCodeFlow(cnt context.Context,rt *token.CodeflowTokenMsg) ( tp *token.Tokenpair, err error)  {
	tp = &token.Tokenpair{&token.AccessTokenMsg{}, &token.RefreshTokenMsg{}}
	tkn, err := tokenanager.ValidateCodeFlow(rt.TokenString)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	cntns, err := storage.CheckTokenStorage(rt.TokenString)
	if err != nil{
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	if !cntns{
		log.Printf("Token not in list:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	tp.AccessToken.TokenString, _, err = tokenanager.NewAccessToken(tkn.UserId, tkn.LogIn)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	var rf int64
	tp.RefreshToken.TokenString, rf, err = tokenanager.NewRefreshToken(tkn.UserId, tkn.LogIn)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	storage.AddTokenStorage(tp.RefreshToken.TokenString, time.Unix(rf,0))
	return
}

func (inst *AuthServerInstance)GetClientInfo(cnt context.Context,id *token.ClientId) (cli *token.ClientInfo, err error){
	cli = &token.ClientInfo{}
	inf, err := models.GetClientInfo(id.Id)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.NotFound, "NotFound ")
		return
	}
	cli.Id = inf.Id
	cli.Name = inf.Name
	cli.Redirurl = inf.RedirUrl
	return
}

