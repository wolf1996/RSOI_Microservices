package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/userclient"
	"net/http"
	"log"
)

func GetUserInfo(c *gin.Context) {
	var inf views.UserInfo
	id := c.MustGet(gin.AuthUserKey).(string)
	res, err := userclient.GetUserInfo(id)
	if err != nil {
		log.Printf("Errpr %s", err.Error())
		err = userclient.ErrorTransform(err)
		code := userclient.StatusCodeFromError(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	inf.UserName = res.Name
	inf.CountEvens = res.Count
	c.JSON(http.StatusOK, inf)
}