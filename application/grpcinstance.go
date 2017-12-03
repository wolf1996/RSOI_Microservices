package application

import (
	"log"
	
	"golang.org/x/net/context"
	"github.com/wolf1996/events/server"
	"github.com/wolf1996/events/application/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
)

type GrpcEventsServerInstance struct {
}

func (inst *GrpcEventsServerInstance)GetEventInfo( cont context.Context,evId *server.EventId) (
	vinf *server.EventInfo,err error){
	inf, err := 	models.GetEventInfo(evId.Id)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Not Found") 
		case models.AddError:
			err = grpc.Errorf(codes.NotFound, "Add Error") 
		default:
			err =  grpc.Errorf(codes.Unknown , "server Error") 
		}
		return
	}
	vinf = &server.EventInfo{
		inf.Id,
		inf.Owner,
		inf.PartCount,
		inf.Description,
	}
	return
}

func (inst *GrpcEventsServerInstance)IncrementUsersNumber(cont context.Context,id *server.EventId) (vinf *server.EventInfo,err error){
	inf, err := models.IncrementEventUserCounter(id.Id)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Not Found") 
		case models.AddError:
			err = grpc.Errorf(codes.NotFound, "Add Error") 
		default:
			err =  grpc.Errorf(codes.Unknown , "server Error") 
		}
		return
	}
	vinf = &server.EventInfo{
	inf.Id,
	inf.Owner,
	inf.PartCount,
	inf.Description,
	}
	return
}
func (inst *GrpcEventsServerInstance)DecrementUsersNumber(cont context.Context,id *server.EventId) (vinf *server.EventInfo,err error){
	inf, err := models.DecrementEventUserCounter(id.Id)
	if err != nil {
		log.Printf(err.Error())
		switch err {
		case models.EmptyResult:
			err = grpc.Errorf(codes.NotFound, "Not Found") 
		case models.AddError:
			err = grpc.Errorf(codes.NotFound, "Add Error") 
		default:
			err =  grpc.Errorf(codes.Unknown , "server Error") 
		}
		return
	}
	vinf = &server.EventInfo{
		inf.Id,
		inf.Owner,
		inf.PartCount,
		inf.Description,
	}
	return
}