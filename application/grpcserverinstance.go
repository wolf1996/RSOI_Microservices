package application

import (
	"github.com/wolf1996/user/server"
	"golang.org/x/net/context"
	"github.com/wolf1996/user/application/models"
)


type UserInfoServer struct{
}

func (srv *UserInfoServer) GetUserInfo(cont context.Context, id *server.UserId) ( infV *server.UserInfo, err error) {
	inf,err := models.GetUserInfo(id.Id)
	if err != nil {
		return
	}
	infV = &server.UserInfo{inf.Name, inf.Count}
	return
}