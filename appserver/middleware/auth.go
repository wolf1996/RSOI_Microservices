package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/resources/authclient"
	"log"
	"github.com/wolf1996/gateway/token"
	"github.com/wolf1996/gateway/appserver/views"
	"net/http"
)

var (
	AtokenName = string("tokenname")
)
func updateTokens(c *gin.Context)(tk token.Token, err error){
	reftoken, err := c.Cookie("RefreshToken")
	if err != nil {
		return
	}
	access, ref, err := authclient.UpdateTokens(reftoken)
	if err != nil {
		return
	}
	c.SetCookie("RefreshToken",ref,0,"","127.0.0.1",false,false)
	c.SetCookie("AccessToken",access,0,"","127.0.0.1",false,false)
	log.Printf("MIDDLEWARE: Token set %s", ref)
	tk, err  = authclient.ValidateAccessToken(access)
	return
}

func processToken(c *gin.Context) (tk token.Token, err error){
	accToken, err := c.Cookie("AccessToken")
	if err != nil {
		return
	}
	tk, err = authclient.ValidateAccessToken(accToken)
	if err != nil {
		tk, err = updateTokens(c)
		if err != nil {
			return
		}
	}
	return
}

func TokenAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tkn, err := processToken(c)
		if err != nil {
			log.Printf("MIDDLEWARE: %s", err.Error())
			c.JSON(http.StatusBadRequest, views.Error{"Token validation error"})
			c.Abort()
			return
		}
		log.Printf("Token set id=%d login=%s", tkn.Id, tkn.LogIn)
		c.Set(AtokenName, tkn)
		c.Next()
		return
	}
}
