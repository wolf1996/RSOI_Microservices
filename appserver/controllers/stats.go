package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/resources/statsclient"
	"log"
	"github.com/wolf1996/gateway/appserver/views"
	"net/http"
	"github.com/wolf1996/gateway/appserver/middleware"
	"github.com/wolf1996/gateway/token"
)

func GetViewsStats(c *gin.Context){
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	logs, err := statsclient.GetViewEvents(tkn)
	if err != nil {
		log.Print(err.Error())
		err, code := statsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	var lst []views.ViewEvent
	for _, i := range logs{
		lst = append(lst, views.ViewEvent{Path:i.Path, UserId:i.UserId})
	}
	c.JSON(http.StatusOK, lst)
}

func GetChangesStats(c *gin.Context){
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	logs, err := statsclient.GetChangeEvents(tkn)
	if err != nil {
		log.Print(err.Error())
		err, code := statsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	var lst []views.ChangeEvent
	for _, i := range logs{
		lst = append(lst, views.ChangeEvent{Path:i.Path, UserId:i.UserId})
	}
	c.JSON(http.StatusOK, lst)
}

func GetLoginStats(c *gin.Context){
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	logs, err := statsclient.GetLogins(tkn)
	if err != nil {
		log.Print(err.Error())
		err, code := statsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	var lst []views.LoginEvent
	for _, i := range logs{
		lst = append(lst, views.LoginEvent{Ok:i.Ok, Info:i.Info})
	}
	c.JSON(http.StatusOK, lst)
}