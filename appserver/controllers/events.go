package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"net/http"
	"strconv"
	"log"
	"github.com/wolf1996/stats/client"
	"github.com/wolf1996/gateway/appserver/middleware"
	"github.com/wolf1996/gateway/token"
)

func GetEventInfo(c *gin.Context) {
	tknBuf, exists := c.Get(middleware.AtokenName)
	var tkn token.Token
	if exists {
		tkn = tknBuf.(token.Token)
	} else {
		tkn.Role = token.Token_ANONIM
	}
	client.WriteInfoViewMessage(c.Request.URL.Path,"")
	var inf views.EventInfo
	key,err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	info, err := eventsclient.GetEventInfo(key, tkn)
	if err != nil {
		log.Print(err.Error())
		err, code := eventsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	inf = views.EventInfo{
		info.Id,
		info.Owner,
		info.PartCount,
		info.Description,
	}
	c.JSON(http.StatusOK, inf)
}

func GetEvents(c *gin.Context){
	tknBuf, exists := c.Get(middleware.AtokenName)
	var tkn token.Token
	if exists {
		tkn = tknBuf.(token.Token)
	} else {
		tkn.Role = token.Token_ANONIM
	}
	client.WriteInfoViewMessage(c.Request.URL.Path,"")
	strparam := c.Param("pagenum")
	if len(strparam) == 0 {
		strparam = "1"
	}
	pnum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	psize := c.Query("pagesize")
	if len(psize) == 0 {
		psize = "1"
	}
	pasize,err := strconv.ParseInt(psize, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	info, err := eventsclient.GetEvents(pasize, pnum, "", tkn)
	if err != nil {
		log.Print(err.Error())
		err, code := eventsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	var infs []views.EventInfo
	for _, i := range info{
		infs = append(infs, views.EventInfo{i.Id, i.Owner, i.PartCount, i.Description})
	}
	c.JSON(http.StatusOK, infs)
}