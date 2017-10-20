package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/appserver/resources"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	var inf views.UserInfo
	id := c.MustGet(gin.AuthUserKey).(string)
	res, err := resources.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	inf.UserName = res.Name
	inf.CountEvens = res.Count
	c.JSON(http.StatusOK, inf)
}