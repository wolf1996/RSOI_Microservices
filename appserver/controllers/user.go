package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/userclient"
	"net/http"
	"log"
	"github.com/gin-gonic/gin/binding"
	"github.com/wolf1996/gateway/resources/authclient"
	"github.com/wolf1996/gateway/appserver/middleware"
	"github.com/wolf1996/gateway/token"
	"github.com/wolf1996/stats/client"
	"strconv"
)

func GetUserInfo(c *gin.Context) {
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	id := tkn.Id
	client.WriteInfoViewMessage(c.Request.URL.Path,id)
	var inf views.UserInfo
	res, err := userclient.GetUserInfo(id, tkn)
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

func GetAccess(c *gin.Context) {
	key,err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	inf, err := authclient.GetClientInfo(key)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	c.JSON(http.StatusOK, views.ClientInfo{inf.Id, inf.Name, inf.RedirUrl})
}

func AllowAccess(c *gin.Context) {
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	urlRed := c.Query("redirect_url")
	if urlRed == ""{
		log.Print("Missed Query parameters")
		c.JSON(http.StatusBadRequest, views.Error{"Missed Query parameters"})
		return
	}
	code, err := authclient.GetCodeGrant(tkn)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	c.JSON(http.StatusOK, views.RedirectInfo{urlRed,code})
}

func GetShiftCodeflow(c *gin.Context) {
	cflow := views.CodeFlowView{}
	err := c.ShouldBindWith(&cflow, binding.JSON)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	acode, rcode, err := authclient.GetCodeflowTokenPair(cflow.CodeFlow)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	c.SetCookie("RefreshToken",rcode,0,"",cflow.Domain,false,false)
	c.SetCookie("AccessToken",acode,0,"",cflow.Domain,false,false)
}