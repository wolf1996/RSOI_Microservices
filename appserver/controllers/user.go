package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/userclient"
	"net/http"
	"log"
	"github.com/gin-gonic/gin/binding"
	"github.com/wolf1996/gateway/resources/authclient"
)

func GetUserInfo(c *gin.Context) {
	var inf views.UserInfo
	id := c.MustGet(gin.AuthUserKey).(string)
	res, err := userclient.GetUserInfo(id)
	if err != nil {
		log.Printf("Error %s", err.Error())
		err, code := userclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	inf.UserName = res.Name
	inf.CountEvens = res.Count
	c.JSON(http.StatusOK, inf)
}

func LogIn(c *gin.Context){
	var logInData views.LogIn
	if err := c.ShouldBindWith(&logInData, binding.JSON); err != nil {
		log.Printf("Error: %s", err.Error())
		c.JSON(http.StatusBadRequest, views.Error{"can't bind body"})
		return
	}
	access, refresh, err := authclient.GetTokenPair(authclient.LogInData{logInData.Login, logInData.Pass})
	if err != nil {
 		log.Printf("Error %s:", err.Error())
 		err, code := authclient.ErrorTransform(err)
 		c.JSON(code, views.Error{err.Error()})
	}
	c.SetCookie("RefreshToken",refresh,0,"","127.0.0.1",false,false)
	c.SetCookie("AccessToken",access,0,"","127.0.0.1",false,false)
}