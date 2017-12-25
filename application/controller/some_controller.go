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
	"github.com/gin-gonic/gin/binding"
	"net/url"
)

func EventsList(c *gin.Context){
	strparam := c.Param("page_num")
	ccs := c.Request.Cookies()
	if len(strparam) == 0 {
		strparam = "1"
	}
	pnum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error(),})
		return
	}
	psize := c.Query("pagesize")
	if len(psize) == 0 {
		psize = "1"
	}
	pasize,err := strconv.ParseInt(psize, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error(),})
		return
	}
	var inf []gatewayview.EventInfo
	inf, req, err := client.GetEvents(pnum, pasize, ccs)
	if err != nil {
		log.Print(err.Error())
		c.HTML(req.StatusCode,"error.html", gin.H{"error": err.Error(),})
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(string(bts[:]))
	for _,i := range req.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(req.StatusCode,"eventslist.html",gin.H{"events":inf, "pagenum":pnum, "pagenumPr":pnum-1, "pagenumN":pnum+1, "pageSize": pasize})
}

func RegistreMe(c *gin.Context){
	var reg view.Registre
	ccs := c.Request.Cookies()
	err := c.ShouldBindWith(&reg, binding.Form)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusForbidden,"error.html", gin.H{"error": err.Error(),})
		return
	}
	log.Print(c.Request)
	resp, err := client.RegistreMe(reg.EventId, ccs)
	if err != nil{
		log.Print(err.Error())
		c.HTML(resp.StatusCode,"error.html", gin.H{"error": err.Error(),})
		return
	}
	for _,i := range resp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.Redirect(http.StatusSeeOther, "/events/1")
}

func GetUserInfo(c *gin.Context) {
	ccs := c.Request.Cookies()
	var inf gatewayview.UserInfo
	inf, req, err := client.UserInfo(ccs)
	if err != nil {
		log.Print(err.Error())
		c.HTML(req.StatusCode,"error.html", gin.H{"error": err.Error(),})
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusServiceUnavailable, "error.html", gin.H{"error":err.Error()})
		return
	}
	log.Print(string(bts[:]))
	for _,i := range req.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(req.StatusCode,"userinfo.html",gin.H{"user":inf})
}

func GetUserRegistrations(c *gin.Context){
	strparam := c.Param("page_num")
	ccs := c.Request.Cookies()
	if len(strparam) == 0 {
		strparam = "1"
	}
	pnum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":err.Error()})
		return
	}
	psize := c.Query("pagesize")
	if len(psize) == 0 {
		psize = "1"
	}
	pasize,err := strconv.ParseInt(psize, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":err.Error()})
		return
	}
	var inf []gatewayview.RegistrationInfo
	inf, req, err := client.GetUserRegistrations(pnum, pasize, ccs)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":err.Error()})
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(string(bts[:]))
	for _,i := range req.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(req.StatusCode,"userregs.html",gin.H{"regs":inf, "pagenum":pnum, "pagenumPr":pnum-1, "pagenumN":pnum+1, "pageSize": pasize})
}

func RemoveReg(c *gin.Context){
	var reg view.RemoveReg
	ccs := c.Request.Cookies()
	err := c.ShouldBindWith(&reg, binding.Form)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html",gin.H{"error":err.Error()})
		return
	}
	log.Print(c.Request)
	resp, err := client.RemoveReg(reg.RegId,ccs)
	if err != nil{
		log.Print(err.Error())
		c.HTML(resp.StatusCode, "error.html", gin.H{"error":err.Error()})
		return
	}
	for _,i := range resp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.Redirect(http.StatusSeeOther, "http://127.0.0.1:8081/user_regs/1")
}

func GetEventInfo(c *gin.Context) {
	var inf gatewayview.EventInfo
	ccs := c.Request.Cookies()
	strparam := c.Param("event")
	if len(strparam) == 0 {
		strparam = "1"
	}
	enum,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":err.Error()})
		return
	}
	inf, req, err := client.EventInfo(enum, ccs)
	if err != nil {
		log.Print(err.Error())
		c.HTML(req.StatusCode, "error.html", gin.H{"error":err.Error()})
		return
	}
	bts, err := json.Marshal(inf)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":err.Error()})
		return
	}
	log.Print(string(bts[:]))
	for _,i := range req.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(req.StatusCode,"eventinfo.html",gin.H{"event":inf})
}

func LogInHandler(c *gin.Context)  {
	lgn := view.LoginForm{}
	err := c.BindWith(&lgn, binding.Form)
	ccs := c.Request.Cookies()
	if err != nil{
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":err.Error()})
		return
	}
	resp, err := client.LogIn(lgn.Login,lgn.Password, ccs)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":"Login Failed"})
		return
	}
	if resp.StatusCode != http.StatusOK{
		log.Print("Login Failed")
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error":"Login Failed"})
		return
	}
	log.Print("Login Ok")
	for _,i := range resp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(http.StatusOK, "error.html", gin.H{"error":"Login Ok"})
	return
}

func LogIn(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html",nil)
}

func BigRedButton(c *gin.Context){
	strparam := c.Param("id")
	ccs := c.Request.Cookies()
	id,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error(),})
		return
	}
	inf, rsp, err := client.GetAccessButton(id, ccs)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error(),})
		return
	}
	for _,i := range rsp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(http.StatusOK, "button.html",gin.H{"id":inf.Id,"name":inf.Name, "redir":inf.RedURL})
}

func GiveAccess(c *gin.Context)  {
	strparam := c.Param("id")
	ccs := c.Request.Cookies()
	id,err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error(),})
		return
	}
	urlRed := c.Query("redirect_url")
	if urlRed == ""{
		log.Print("Missed Query parameters")
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error":"Missed Qery parameters",})
		return
	}
	inf, rsp, err := client.GivAccess(id, urlRed, ccs)
	if err != nil {
		log.Print(err.Error())
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": err.Error(),})
		return
	}
	if rsp.StatusCode != http.StatusOK {
		log.Print("Error: failed to parse client url %s")
		c.HTML(http.StatusBadRequest,"error.html", gin.H{"error": "Some error",})
		return
	}
	for _,i := range rsp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	log.Printf("red url %s", inf)
	ur, err := url.Parse(inf.RedirectUrl)
	if rsp.StatusCode != http.StatusOK {
		log.Print("Error: failed to parse client url %s", err.Error())
		c.HTML(http.StatusInternalServerError,"error.html", gin.H{"error": "Some error",})
		return
	}
	q := ur.Query()
	q.Add("code", inf.CodeFlow)
	ur.RawQuery = q.Encode()
	c.Redirect(http.StatusSeeOther, ur.String())
}