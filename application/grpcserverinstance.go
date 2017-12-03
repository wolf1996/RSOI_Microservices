package application

import (
	"log"
	"github.com/wolf1996/user/server"
	"golang.org/x/net/context"
	"github.com/wolf1996/user/application/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
)


type UserInfoServer struct{
}

func (srv *UserInfoServer) GetUserInfo(cont context.Context, id *server.UserId) ( infV *server.UserInfo, err error) {
	inf,err := models.GetUserInfo(id.Id)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Not found") 
		default:
			err =  grpc.Errorf(codes.Unknown , "server Error") 
		}
		return
	}
	infV = &server.UserInfo{inf.Name, inf.Count, inf.Id}
	return
}

func (srv *UserInfoServer) IncrementUserCounter(cont context.Context, id *server.UserId) (infV *server.UserInfo, err error) {
	inf,err := models.IncrementUserEventCounter(id.Id)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Not Found") 
		case nil:
			break
		default:
			err =  grpc.Errorf(codes.Unknown , "server Error") 
		}
	}
	infV = &server.UserInfo{inf.Name, inf.Count, inf.Id}
	return
}

func (srv *UserInfoServer) DecrementUserCounter(cont context.Context, id *server.UserId) (infV *server.UserInfo, err error) {
	inf,err := models.DecrementUserEventCounter(id.Id)
	if err != nil {
		log.Print(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Not Found") 
		default:
			err =  grpc.Errorf(codes.Unknown , "server Error") 
		}
	}
	infV = &server.UserInfo{inf.Name, inf.Count, inf.Id}
	return
}