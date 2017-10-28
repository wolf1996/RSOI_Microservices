package application

import (
	"github.com/wolf1996/registration/server"
	"golang.org/x/net/context"
	"github.com/wolf1996/registration/application/models"
)


type GprsServerInstance struct {
}

func (inst GprsServerInstance)GetRegistrationInfo(cont context.Context,id *server.RegistrationId) ( infV *server.RegistrationInfo, err error) {
	inf,err := models.GetRegistration(id.Id)
	if err != nil {
		return
	}
	infV = &server.RegistrationInfo{inf.Id,inf.UserId, inf.EventId}
	return
}

