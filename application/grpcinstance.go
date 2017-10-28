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
	print("Used")
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