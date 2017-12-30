package application

import (
	"github.com/wolf1996/stats/server"
	"context"
	"github.com/wolf1996/stats/shared"
)

type GrpcServerInstance struct {

}




func (inst *GrpcServerInstance)GetLogins(context.Context, *server.Empty) (res *server.LoginEvents,err error) {
	res = &server.LoginEvents{}
	var msgs []shared.LoginMsg
	err = mgoDb.C(shared.TypeLogin).Find(nil).All(&msgs)
	if err != nil {
		return
	}
	for _, i := range msgs{
		res.Data = append(res.Data, &server.LoginEvent{i.Ok, i.Info})
	}
	return
}
func (inst *GrpcServerInstance)GetChanges(context.Context, *server.Empty) (res *server.ChangeInfos,err error) {
	res = &server.ChangeInfos{}
	var msgs []shared.InfoChangeMsg
	err = mgoDb.C(shared.TypeChangeInf).Find(nil).All(&msgs)
	if err != nil {
		return
	}
	for _, i := range msgs{
		res.Data = append(res.Data, &server.ChangeInfo{i.UserId, i.Path})
	}
	return
}
func(inst *GrpcServerInstance) GetViews(context.Context, *server.Empty) (res *server.ViewInfos,err error) {
	res = &server.ViewInfos{}
	var msgs []shared.InfoViewMsg
	err = mgoDb.C(shared.TypeGetInf).Find(nil).All(&msgs)
	if err != nil {
		return
	}
	for _, i := range msgs{
		res.Data = append(res.Data, &server.ViewInfo{i.UserId, i.Path})
	}
	return
}
