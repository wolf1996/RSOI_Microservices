package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/resources/statsclient"
	"log"
	"github.com/wolf1996/gateway/appserver/views"
	"net/http"
)

func GetViewsStats(c *gin.Context){
	logs, err := statsclient.GetViewEvents()
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
	logs, err := statsclient.GetChangeEvents()
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
	logs, err := statsclient.GetLogins()
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