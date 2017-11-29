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
		err = eventsclient.ErrorTransform(err)
		code := eventsclient.StatusCodeFromError(err)
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