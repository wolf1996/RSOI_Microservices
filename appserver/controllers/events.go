package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"net/http"
	"strconv"
	"log"
)

func GetEventInfo(c *gin.Context) {
	var inf views.EventInfo
	key,err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	info, err := eventsclient.GetEventInfo(key)
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
	info, err := eventsclient.GetEvents(pasize, pnum, "")
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