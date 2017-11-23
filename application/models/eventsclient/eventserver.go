package eventsclient

import (
	"github.com/wolf1996/events/server"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
)

type Config struct{
	Addres string
}

type EventInfo struct {
	Id              int64
	Owner           string
	PartCount       int64
	Description     string
}

var addres string

func SetConfigs(config Config){
	addres = config.Addres
	log.Print(fmt.Sprintf("used to eventInfo service %s", addres))

}


func GetEventInfo(id int64) (uinf *EventInfo,err  error){
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewEventServiceClient(conn)
	info, err := cli.GetEventInfo(context.Background(), &server.EventId{id})
	if err != nil {
		return
	}
	return &EventInfo{info.Id, info.Name, info.Participants, info.Description}, nil
}

func IncrementEventUsers(id int64) (uinf *EventInfo,err  error) {
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewEventServiceClient(conn)
	info, err := cli.IncrementUsersNumber(context.Background(), &server.EventId{id})
	if err != nil {
		return
	}
	return &EventInfo{info.Id, info.Name, info.Participants, info.Description}, nil
}

func DecrementEventUsers(id int64) (uinf *EventInfo,err  error) {
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		return
	}
	cli := server.NewEventServiceClient(conn)
	info, err := cli.DecrementUsersNumber(context.Background(), &server.EventId{id})
	if err != nil {
		return
	}
	return &EventInfo{info.Id, info.Name, info.Participants, info.Description}, nil
}