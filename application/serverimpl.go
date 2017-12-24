package application

import (
	"github.com/wolf1996/auth/token"
	"github.com/wolf1996/auth/application/tokenanager"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/wolf1996/auth/application/models"
	"log"
)

type AuthServerInstance struct {
}
/*
	GetAccessToken(context.Context, *RefreshTokenMsg) (*AccessTokenMsg, error)
	GetTokenpair(context.Context, *SignInPair) (*Tokenpair, error)
	RefreshTokenValidation(context.Context, *RefreshTokenMsg) (*ValidResult, error)
	AccessTokenValidation(context.Context, *AccessTokenMsg) (*ValidResult, error)
 */
func (inst *AuthServerInstance)GetAccessToken(cnt context.Context, rfrsh *token.RefreshTokenMsg) (msg *token.Tokenpair,err error){
	tkn, err := tokenanager.ValidateRefreshToken(rfrsh.TokenString)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
	}
	actoken, err := tokenanager.RefreshAccessToken(tkn)
	restoken, err := tokenanager.RefreshRefreshToken(tkn)
	msg.RefreshToken = &token.RefreshTokenMsg{ restoken}
	msg.AccessToken  = &token.AccessTokenMsg{ actoken}
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
	tkns.AccessToken.TokenString, err = tokenanager.ProduceAccessToken(uinf)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	tkns.RefreshToken.TokenString, err = tokenanager.ProduceRefreshToken(uinf)
	if err != nil {
		log.Printf("ERROR:%s", err.Error())
		err = status.Errorf(codes.InvalidArgument, "Token validation failed")
		return
	}
	return
}

func (inst *AuthServerInstance)AccessTokenValidation(cnt context.Context,at *token.AccessTokenMsg) (val *token.ValidResult,err error){
	val = &token.ValidResult{Valid:false}
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