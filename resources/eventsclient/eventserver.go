package eventsclient

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/credentials"

	"github.com/wolf1996/gateway/evserver"
	"google.golang.org/grpc"
	"github.com/wolf1996/gateway/token"
)

type Config struct {
	Addres           string
	Crt              string
	QuemanagerConfig QConfig
}

type EventInfo struct {
	Id          int64
	Owner       string
	PartCount   int64
	Description string
}

var addres string
var ConnectionError = fmt.Errorf("Can't connect to Events")
var creds credentials.TransportCredentials

func SetConfigs(config Config) {
	addres = config.Addres
	log.Print(fmt.Sprintf("used to eventInfo service %s", addres))
	err := ApplyConfig(config.QuemanagerConfig)
	if err != nil {
		panic(err.Error())
	}
	creds, err = credentials.NewClientTLSFromFile(config.Crt, "")
	if err != nil {
		panic(err.Error())
	}
}

func GetEventInfo(id int64, token token.Token) (uinf *EventInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := evserver.NewEventServiceClient(conn)
	info, err := cli.GetEventInfo(context.Background(), &evserver.EventId{id})
	if err != nil {
		return
	}
	return &EventInfo{info.Id, info.Name, info.Participants, info.Description}, nil
}

func IncrementEventUsers(id int64, token token.Token) (uinf *EventInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := evserver.NewEventServiceClient(conn)
	info, err := cli.IncrementUsersNumber(context.Background(), &evserver.EventId{id})
	if err != nil {
		return
	}
	return &EventInfo{info.Id, info.Name, info.Participants, info.Description}, nil
}

func DecrementEventUsers(id int64, token token.Token) (uinf *EventInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	cli := evserver.NewEventServiceClient(conn)
	info, err := cli.DecrementUsersNumber(context.Background(), &evserver.EventId{id})
	if err != nil {
		return
	}
	return &EventInfo{info.Id, info.Name, info.Participants, info.Description}, nil
}

func DecrementEventUsersAsync(id int64, token token.Token) (err error) {
	DecrementEventsCounter(id, token)
	return
}

func GetEvents(pageSize int64, pageNum int64, userId string, token token.Token) (events []EventInfo, err error) {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(creds))
	if err != nil {
		err = ConnectionError
		return
	}
	cli := evserver.NewEventServiceClient(conn)
	evStream, err := cli.GetEvents(context.Background(), &evserver.EventsRequest{pageSize, pageNum, userId})
	if err != nil {
		log.Print(err)
		return
	}
	var inf *evserver.EventInfo
	for {
		inf, err = evStream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Print(err.Error())
				return
			}
			err = nil
			evStream.CloseSend()
			return
		}
		events = append(events, EventInfo{inf.Id, inf.Name, inf.Participants, inf.Description})
	}
	return
}
