package tokenanager

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaime struct{
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	LogIn  string `json:"log_in"`
}

type RefreshTokenClaime struct{
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	LogIn  string `json:"log_in"`
}

func (tkn *AccessTokenClaime)Valid() error {
	return nil
}


func (tkn *RefreshTokenClaime)Valid() error {
	return nil
}