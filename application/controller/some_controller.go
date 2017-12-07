package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/frontend/application/client"
	"github.com/wolf1996/frontend/application/client/gatewayview"
	"log"
	"encoding/json"
	"strconv"
	"net/http"
	"github.com/wolf1996/frontend/application/view"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/gin-gonic/gin/binding"
)

func EventsList(c *gin.Context){
	strparam := c.Param("page_num")
	if len(strparam) == 0 {
		strparam = "1"
	}
	pnum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
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
	var inf []gatewayview.EventInfo
	inf, req, err := client.GetEvents(pnum, pasize)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print(string(bts[:]))
	c.HTML(req.StatusCode,"eventslist.html",gin.H{"events":inf, "pagenum":pnum, "pagenumPr":pnum-1, "pagenumN":pnum+1, "pageSize": pasize})
}

func RegistreMe(c *gin.Context){
	var reg view.Registre
	err := c.ShouldBindWith(&reg, binding.Form)
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print(c.Request)
	resp, err := client.RegistreMe(reg.EventId)
	if err != nil{
		log.Print(err.Error())
		c.Status(resp.StatusCode)
		return
	}
	c.Redirect(http.StatusSeeOther, "http://127.0.0.1:8081/events/1")
}

func GetUserInfo(c *gin.Context) {
	var inf gatewayview.UserInfo
	inf, req, err := client.UserInfo()
	if err != nil {
		log.Print(err.Error())
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print(string(bts[:]))
	c.HTML(req.StatusCode,"userinfo.html",gin.H{"user":inf})
}

func GetUserRegistrations(c *gin.Context){
	strparam := c.Param("page_num")
	if len(strparam) == 0 {
		strparam = "1"
	}
	pnum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
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
	}
	var inf []gatewayview.RegistrationInfo
	inf, req, err := client.GetUserRegistrations(pnum, pasize)
	if err != nil {
		log.Print(err.Error())
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(string(bts[:]))
	c.HTML(req.StatusCode,"userregs.html",gin.H{"regs":inf, "pagenum":pnum, "pagenumPr":pnum-1, "pagenumN":pnum+1, "pageSize": pasize})
}

func RemoveReg(c *gin.Context){
	var reg view.RemoveReg
	err := c.ShouldBindWith(&reg, binding.Form)
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print(c.Request)
	resp, err := client.RemoveReg(reg.RegId)
	if err != nil{
		log.Print(err.Error())
		c.JSON(resp.StatusCode, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "http://127.0.0.1:8081/user_regs/1")
}

func GetEventInfo(c *gin.Context) {
	var inf gatewayview.EventInfo
	strparam := c.Param("event")
	if len(strparam) == 0 {
		strparam = "1"
	}
	enum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		return
	}
	inf, req, err := client.EventInfo(enum)
	if err != nil {
		log.Print(err.Error())
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print(string(bts[:]))
	c.HTML(req.StatusCode,"eventinfo.html",gin.H{"event":inf})
}