package application

import (
	"github.com/wolf1996/user/server"
	"golang.org/x/net/context"
	"fmt"

)

type UserInfo struct {
	Name string
	Count int64
}

var userMap = map [string]UserInfo {
	"simpleUser": UserInfo{"Ivanov invan ivanovich", 4},
	"eventOwner": UserInfo{"Kakoi-to chuvak", 3},
}


type UserInfoServer struct{
}

func (srv *UserInfoServer) GetUserInfo(cont context.Context, id *server.UserId) (*server.UserInfo, error) {
	inf, ok := userMap[id.Id]
	if !ok{
		return nil, fmt.Errorf("Can't find user by id %s", id)
	}
	return &server.UserInfo{inf.Name, inf.Count}, nil
}