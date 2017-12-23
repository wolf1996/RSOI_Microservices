package tokenanager

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
	"github.com/wolf1996/auth/application/models"
)

type Config struct {
	Salt       string
	AccessExp  int64
	RefreshExp int64
}


var (
	salt       []byte
	accessExp  int64
	refreshExp int64
)


func ApplyConfig(config Config){
	salt = []byte(config.Salt)
	accessExp = config.AccessExp
	refreshExp = config.RefreshExp
}

func ValidateAccessToken(token string)(tk AccessTokenClaime, err error){
	parsedToken, err := jwt.ParseWithClaims(token, &AccessTokenClaime{},func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return salt, nil
	})
	if claims,ok  := parsedToken.Claims.(*AccessTokenClaime); ok && parsedToken.Valid{
		tk = *claims
	}
	return
}


func ValidateRefreshToken(token string)(tk RefreshTokenClaime, err error){
	parsedToken, err := jwt.ParseWithClaims(token, &RefreshTokenClaime{},func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return salt, nil
	})
	if claims,ok  := parsedToken.Claims.(*RefreshTokenClaime); ok && parsedToken.Valid{
		tk = *claims
	}
	return
}

func RefreshRefreshToken(tk RefreshTokenClaime) (tkn string, err error){
	newtk := RefreshTokenClaime{UserId:tk.UserId}
	newtk.ExpiresAt = time.Now().Unix() + refreshExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &newtk)
	tkn, err = token.SignedString(salt)
	return
}

func RefreshAccessToken(tk RefreshTokenClaime) (tkn string, err error){
	newtk := AccessTokenClaime{UserId:tk.UserId}
	newtk.ExpiresAt = time.Now().Unix() + accessExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &newtk)
	tkn, err = token.SignedString(salt)
	return
}

func ProduceRefreshToken(info models.UserInfo) (tkn string, err error){
	newtk := RefreshTokenClaime{UserId:info.Id}
	newtk.ExpiresAt = time.Now().Unix() + refreshExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &newtk)
	tkn, err = token.SignedString(salt)
	return
}

func ProduceAccessToken(info models.UserInfo) (tkn string, err error){
	newtk := AccessTokenClaime{UserId:info.Id}
	newtk.ExpiresAt = time.Now().Unix() + accessExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &newtk)
	tkn, err = token.SignedString(salt)
	return
}