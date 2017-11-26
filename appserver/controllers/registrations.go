package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/registrationclient"
	"log"
	"net/http"
	"strconv"
	"github.com/wolf1996/gateway/resources/userclient"
	_ "github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"
	_ "github.com/wolf1996/gateway/authtoken"
	"github.com/golang/protobuf/proto"
	"github.com/wolf1996/gateway/authtoken"
	"encoding/base64"
	"github.com/wolf1996/gateway/resources/eventsclient"
)

func RegistrateMe(c *gin.Context) {
	//добавить токен здесь
	user := c.MustGet(gin.AuthUserKey).(string)
	key,err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	eventData, err := eventsclient.IncrementEventUsers(key)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	defer func(){
		if err != nil {
			log.Print("Some error occured, revert eventsclient counter")
			_, errDef := eventsclient.DecrementEventUsers(key)
			if errDef != nil{
				log.Print("Defer error")
			}
		}
	}()
	userData, err := userclient.IncrementEventsCounter(user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}

	defer func(){
		if err != nil {
			log.Print("Some error occured, revert user counter")
			_, errDef := userclient.DecrementEventsCounter(user)
			if errDef != nil{
				log.Print("Defer error")
			}
		}
	}()

	regdata, err := registrationclient.AddRegistration(userData.Name, eventData.Id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	res := views.AllRegInfo{regdata.Id,
		views.EventInfo{eventData.Id,
				eventData.Owner,
				eventData.PartCount,
				 eventData.Description,
			},
			views.UserInfo{
				userData.Name,
				userData.Count,
				userData.Id,
			},
	}
	c.JSON(http.StatusOK, res)
}


func RemoveRegistration(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	key,err := strconv.ParseInt(c.Param("registration_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	token := authtoken.Token{user}
	btTok,err := proto.Marshal(&token)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	strTok := base64.StdEncoding.EncodeToString(btTok)
	md := metadata.Pairs("token", strTok)
	regdata, err := registrationclient.RemoveRegistration(key, md)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	err = eventsclient.DecrementEventUsersAsync(regdata.EventId)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	userData, err := userclient.DecrementEventsCounter(user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	res := views.AllRegInfo2{regdata.Id,
		regdata.EventId,
		views.UserInfo{
			userData.Name,
			userData.Count,
			userData.Id,
		},
	}
	c.JSON(http.StatusOK, res)
}

func GetRegisrationInfo(c *gin.Context){
	var inf views.RegistrationInfo
	key,err := strconv.ParseInt(c.Param("registration_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	info, err := registrationclient.GetRegistrationInfo(key)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	inf = views.RegistrationInfo{
		info.Id,
		info.UserId,
		info.EventId,
	}
	c.JSON(http.StatusOK, inf)
}

func GetRegistrations(c *gin.Context){
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
	var infs []views.RegistrationInfo
	id := c.MustGet(gin.AuthUserKey).(string)
	res, err := registrationclient.GetRegistrations(id,pnum,1)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	for _, i := range res{
		infs = append(infs, views.RegistrationInfo{i.Id, i.UserId, i.EventId})
	}
	c.JSON(http.StatusOK, infs)
}