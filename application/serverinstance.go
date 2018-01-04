package application

import (
	"github.com/wolf1996/stats/server"
	"github.com/wolf1996/stats/application/model"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"log"
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"
	"github.com/wolf1996/stats/token"
	"context"
	"google.golang.org/grpc"
)

type GrpcServerInstance struct {

}

/*
	GetLogins(*Empty, StatisticService_GetLoginsServer) error
	GetChanges(*Empty, StatisticService_GetChangesServer) error
	GetViews(*Empty, StatisticService_GetViewsServer) error
 */

func decodeTokenString(bt string)(tkn token.Token,err error){
	bts, err := base64.StdEncoding.DecodeString(bt)
	if err != nil {
		return
	}
	err = proto.Unmarshal(bts, &tkn)
	return
}

func getTokenFromContext(cont context.Context)(tkn token.Token, err error){
	md, ok := metadata.FromIncomingContext(cont)
	if !ok {
		log.Print("Can't find metadata")
		err = grpc.Errorf(codes.InvalidArgument, "Can't find metadata")
		return
	}
	tks, ok := md["token"]
	if (!ok) || (len(tks) < 1) {
		log.Print("Can't find token")
		err = grpc.Errorf(codes.InvalidArgument, "Can't find token")
		return
	}
	return decodeTokenString(tks[0])
}


func (inst *GrpcServerInstance)GetLogins(e *server.Empty,resStream server.StatisticService_GetLoginsServer) (err error) {
	tkn, err := getTokenFromContext(resStream.Context())
	if err != nil {
		return
	}
	if tkn.Role != token.Token_ADMIN {
		log.Printf("Failed to access %v", tkn)
		return status.Errorf(codes.PermissionDenied, "You need to be admin to do it")
	}
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
	tkn, err := getTokenFromContext(resStream.Context())
	if err != nil {
		return
	}
	if tkn.Role != token.Token_ADMIN {
		log.Printf("Failed to access %v", tkn)
		return status.Errorf(codes.PermissionDenied, "You need to be admin to do it")
	}
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
	tkn, err := getTokenFromContext(resStream.Context())
	if err != nil {
		return
	}
	if tkn.Role != token.Token_ADMIN {
		log.Printf("Failed to access %v", tkn)
		return status.Errorf(codes.PermissionDenied, "You need to be admin to do it")
	}
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
