package controllers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/proto"
	"github.com/wolf1996/gateway/appserver/middleware"
	"github.com/wolf1996/gateway/appserver/views"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"github.com/wolf1996/gateway/resources/registrationclient"
	"github.com/wolf1996/gateway/resources/userclient"
	"github.com/wolf1996/gateway/token"
	"github.com/wolf1996/stats/client"
	"google.golang.org/grpc/metadata"
)

func RegistrateMe(c *gin.Context) {
	//добавить токен здесь
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	user := tkn.LogIn
	client.WriteInfoChangeMessage(c.Request.URL.Path, user)
	key, err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	eventData, err := eventsclient.IncrementEventUsers(key)
	if err != nil {
		log.Print(err.Error())
		err, code := eventsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	defer func() {
		if err != nil {
			log.Print("Some error occured, revert eventsclient counter")
			_, errDef := eventsclient.DecrementEventUsers(key)
			if errDef != nil {
				log.Print("Defer error")
			}
		}
	}()
	userData, err := userclient.IncrementEventsCounter(user)
	if err != nil {
		log.Print(err.Error())
		err, code := userclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}

	defer func() {
		if err != nil {
			log.Print("Some error occured, revert user counter")
			_, errDef := userclient.DecrementEventsCounter(user)
			if errDef != nil {
				log.Print("Defer error")
			}
		}
	}()

	regdata, err := registrationclient.AddRegistration(userData.Name, eventData.Id)
	if err != nil {
		log.Print(err.Error())
		err, code := registrationclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
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
	c.JSON(http.StatusCreated, res)
}

func RemoveRegistration(c *gin.Context) {
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	user := tkn.LogIn
	client.WriteInfoChangeMessage(c.Request.URL.Path, user)
	key, err := strconv.ParseInt(c.Param("registration_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	btTok, err := proto.Marshal(&tkn)
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
		err, code := registrationclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	err = eventsclient.DecrementEventUsersAsync(regdata.EventId)
	if err != nil {
		log.Print(err.Error())
		err, code := eventsclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	err = userclient.DecrementEventsCounterAsync(user)
	if err != nil {
		log.Print(err.Error())
		err, code := userclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	res := views.AllRegInfoAsync{regdata.Id,
		regdata.EventId,
		regdata.UserId,
	}
	c.JSON(http.StatusOK, res)
}

func GetRegisrationInfo(c *gin.Context) {
	var inf views.RegistrationInfo
	client.WriteInfoViewMessage(c.Request.URL.Path, "")
	key, err := strconv.ParseInt(c.Param("registration_id"), 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusNotFound, views.Error{err.Error()})
		return
	}
	info, err := registrationclient.GetRegistrationInfo(key)
	if err != nil {
		log.Print(err.Error())
		err, code := registrationclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	inf = views.RegistrationInfo{
		info.Id,
		info.UserId,
		info.EventId,
	}
	c.JSON(http.StatusOK, inf)
}

func GetRegistrations(c *gin.Context) {
	tkn := c.MustGet(middleware.AtokenName).(token.Token)
	id := tkn.LogIn
	client.WriteInfoViewMessage(c.Request.URL.Path, id)
	strparam := c.Param("pagenum")
	if len(strparam) == 0 {
		strparam = "1"
	}
	pnum, err := strconv.ParseInt(strparam, 10, 64)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, views.Error{err.Error()})
		return
	}
	var infs []views.RegistrationInfo
	res, err := registrationclient.GetRegistrations(id, pnum, 1)
	if err != nil {
		log.Print(err.Error())
		err, code := registrationclient.ErrorTransform(err)
		c.JSON(code, views.Error{err.Error()})
		return
	}
	for _, i := range res {
		infs = append(infs, views.RegistrationInfo{i.Id, i.UserId, i.EventId})
	}
	c.JSON(http.StatusOK, infs)
}
