package application

import (
	"github.com/wolf1996/stats/server"
	"github.com/wolf1996/stats/application/model"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"log"
)

type GrpcServerInstance struct {

}

/*
	GetLogins(*Empty, StatisticService_GetLoginsServer) error
	GetChanges(*Empty, StatisticService_GetChangesServer) error
	GetViews(*Empty, StatisticService_GetViewsServer) error
 */

func (inst *GrpcServerInstance)GetLogins(e *server.Empty,resStream server.StatisticService_GetLoginsServer) (err error) {
	msgs, err  := model.GetLogins()
	if err != nil{
		log.Printf("ERROR get logins info: %s", err.Error())
		err = status.Errorf(codes.Unavailable, "some error occured")
		return
	}
	for _, i := range msgs{
		resStream.Send(&server.LoginEvent{Ok:i.Ok, Result:i.Info})
	}
	return
}
func (inst *GrpcServerInstance)GetChanges(e *server.Empty,resStream server.StatisticService_GetChangesServer) (err error) {
	msgs, err := model.GetChangeEvents()
	if err != nil{
		log.Printf("ERROR get changes info : %s", err.Error())
		err = status.Errorf(codes.Unavailable, "some error occured")
		return
	}
	for _, i := range msgs{
		resStream.Send(&server.ChangeInfo{UserLogin:i.UserId, Path:i.Path})
	}
	return
}
func(inst *GrpcServerInstance) GetViews(e *server.Empty,resStream server.StatisticService_GetViewsServer) (err error) {
	msgs, err := model.GetViewEvents()
	if err != nil{
		log.Printf("ERROR get views info : %s", err.Error())
		err = status.Errorf(codes.Unavailable, "some error occured")
		return
	}
	for _, i := range msgs{
		resStream.Send(&server.ViewInfo{i.UserId, i.Path})
	}
	return
}
