package application

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/wolf1996/registration/application/models"
	"github.com/wolf1996/registration/server"
	"github.com/wolf1996/registration/token"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"encoding/base64"
)

type GprsServerInstance struct {
}


func decodeTokenString(bt string)(tkn token.Token,err error){
	tk := token.Token{}
	bts, err := base64.StdEncoding.DecodeString(bt)
	if err != nil {
		return
	}
	err = proto.Unmarshal(bts, &tk)
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

func (inst GprsServerInstance) GetRegistrationInfo(cont context.Context, id *server.RegistrationId) (infV *server.RegistrationInfo, err error) {
	infV = new(server.RegistrationInfo)
	inf, err := models.GetRegistration(id.Id)
	if err != nil {
		log.Print(err.Error())
	}
	switch err {
	case models.EmptyResult:
		err = grpc.Errorf(codes.NotFound, "Can't find any registrations")
	case nil:
		break
	default:
		err = grpc.Errorf(codes.Unknown, "server Error")
	}
	infV = &server.RegistrationInfo{inf.Id, inf.UserId, inf.EventId}
	return
}

func (inst GprsServerInstance) AddRegistration(ctx context.Context, in *server.RegistrationToAdd) (infV *server.RegistrationInfo, err error) {
	infV = new(server.RegistrationInfo)
	inf, err := models.AddRegistration(in.UserId, in.EventId)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Empty result")
		case models.AddError:
			err = grpc.Errorf(codes.NotFound, "Can't add registration")
		default:
			err = grpc.Errorf(codes.Unknown, "server Error")
		}
		return
	}
	infV = &server.RegistrationInfo{inf.Id, inf.UserId, inf.EventId}
	return
}

func (inst GprsServerInstance) RemoveRegistration(cont context.Context, id *server.RegistrationId) (infV *server.RegistrationInfo, err error) {
	infV = new(server.RegistrationInfo)
	_ , err = getTokenFromContext(cont)
	if err != nil {
		return
	}
	if err != nil {
		err = grpc.Errorf(codes.InvalidArgument, "Can't parse argument")
		log.Print("Can't parse argument %s", err.Error())
		return
	}
	inf, err := models.RemoveRegistration(id.Id)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Empty result")
		case models.AddError:
			err = grpc.Errorf(codes.NotFound, "Can't add registration")
		default:
			err = grpc.Errorf(codes.Unknown, "server Error")
		}
		return
	}
	infV = &server.RegistrationInfo{inf.Id, inf.UserId, inf.EventId}
	return
}

func (inst GprsServerInstance) GetUserRegistrations(req *server.UsersRegistrationsRequest, stream server.RegistrationService_GetUserRegistrationsServer) (err error) {
	if req.PageNumber <= 0 {
		return grpc.Errorf(codes.InvalidArgument, "Invalid page name %s", req.PageSize)
	}
	if req.PageSize <= 0 {
		return grpc.Errorf(codes.InvalidArgument, "Invalid page size %s", req.PageSize)
	}
	err = models.GetUserRegistrations(req.UserId, req.PageNumber-1, req.PageSize, stream)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Empty result")
		case models.AddError:
			err = grpc.Errorf(codes.NotFound, "Can't add registration")
		default:
			err = grpc.Errorf(codes.Unknown, "server Error")
		}
		return
	}
	return
}
