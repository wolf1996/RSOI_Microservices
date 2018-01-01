package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/frontend/application/client"
	"log"
	"net/http"
)

func GetViewsStats(c *gin.Context){
	ccs := c.Request.Cookies()
	views, resp, err := client.GetViewStats(ccs)
	if err != nil {
		log.Print(err.Error())
		var code int
		if resp != nil {
			code = resp.StatusCode
		} else {
			code =  http.StatusServiceUnavailable
		}
		c.HTML(code,"error.html", gin.H{"error": err.Error(),})
		return
	}
	for _,i := range resp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(http.StatusOK, "views_stats.html", gin.H{"events":views,})
}

func GetChangesStats(c *gin.Context){
	ccs := c.Request.Cookies()
	views, resp, err := client.GetChangeStats(ccs)
	if err != nil {
		log.Print(err.Error())
		var code int
		if resp != nil {
			code = resp.StatusCode
		} else {
			code =  http.StatusServiceUnavailable
		}
		c.HTML(code,"error.html", gin.H{"error": err.Error(),})
		return
	}
	for _,i := range resp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(http.StatusOK, "change_stats.html", gin.H{"events":views,})
}

func GetLoginStats(c *gin.Context){
	ccs := c.Request.Cookies()
	views, resp, err := client.GetLoginStats(ccs)
	if err != nil {
		log.Print(err.Error())
		var code int
		if resp != nil {
			code = resp.StatusCode
		} else {
			code =  http.StatusServiceUnavailable
		}
		c.HTML(code,"error.html", gin.H{"error": err.Error(),})
		return
	}
	for _,i := range resp.Cookies(){
		c.SetCookie(i.Name,i.Value,i.MaxAge,i.Path,i.Domain,i.Secure, i.HttpOnly)
	}
	c.HTML(http.StatusOK, "login_stats.html", gin.H{"events":views,})
}