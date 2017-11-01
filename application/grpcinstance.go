package application

import (
	"golang.org/x/net/context"
	"github.com/wolf1996/events/server"
	"github.com/wolf1996/events/application/models"
)

type GrpcEventsServerInstance struct {
}

func (inst *GrpcEventsServerInstance)GetEventInfo( cont context.Context,evId *server.EventId) (
	vinf *server.EventInfo,err error){
	inf, err := 	models.GetEventInfo(evId.Id)
	if err != nil {
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
