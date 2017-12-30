package tokenanager

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaime struct{
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	LogIn  string `json:"log_in"`
	Role   int32 `json:"role"`
}

type RefreshTokenClaime struct{
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	LogIn  string `json:"log_in"`
	Role   int32 `json:"role"`
}

type CodeFlowClaime struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	LogIn  string `json:"log_in"`
}

func (tkn *CodeFlowClaime)Valid() error {
	return tkn.StandardClaims.Valid()
}

func (tkn *AccessTokenClaime)Valid() error {
	return tkn.StandardClaims.Valid()
}


func (tkn *RefreshTokenClaime)Valid() error {
	return tkn.StandardClaims.Valid()
}