package application

import (
	"github.com/wolf1996/registration/server"
	"golang.org/x/net/context"
	"github.com/wolf1996/registration/application/models"
	"google.golang.org/grpc/metadata"
	"github.com/golang/protobuf/proto"
	"fmt"
	"log"
	"github.com/wolf1996/gateway/authtoken"
	"encoding/base64"
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

func (inst GprsServerInstance)AddRegistration(ctx context.Context, in *server.RegistrationToAdd) ( infV *server.RegistrationInfo, err error) {
	inf,err := models.AddRegistration(in.UserId, in.EventId)
	if err != nil {
		return
	}
	infV = &server.RegistrationInfo{inf.Id,inf.UserId, inf.EventId}
	return
}

func (inst GprsServerInstance)RemoveRegistration(cont context.Context,id *server.RegistrationId) ( infV *server.RegistrationInfo, err error) {
	md, ok := metadata.FromIncomingContext(cont)
	if !ok{
		log.Print("Can't find metadata")
		err = fmt.Errorf("Can't find metadata")
		return
	}
	tks,ok := md["token"]
	if (!ok) || (len(tks) < 1){
		log.Print("Can't find token")
		err = fmt.Errorf("Can't find token")
		return
	}
	tk := authtoken.Token{}
	btk,err := base64.StdEncoding.DecodeString(tks[0])
	err = proto.Unmarshal(btk, &tk)
	if err != nil {
		return
	}
	log.Printf("User Id = %s", tk.Id)
	inf,err := models.RemoveRegistration(id.Id)
	if err != nil {
		return
	}
	infV = &server.RegistrationInfo{inf.Id,inf.UserId, inf.EventId}
	return
}

func (inst GprsServerInstance)	GetUserRegistrations(req *server.UsersRegistrationsRequest, stream server.RegistrationService_GetUserRegistrationsServer) (err error) {
	if req.PageNumber <= 0 {
		return fmt.Errorf("Invalid page name %s", req.PageNumber)
	}
	if req.PageSize <= 0 {
		return fmt.Errorf("Invalid page size %s", req.PageSize)
	}
	err = models.GetUserRegistrations(req.UserId, req.PageNumber - 1, req.PageSize, stream )
	return
}