package eventsclient

import (
	"github.com/wolf1996/gateway/resources"
	"github.com/streadway/amqp"
	"fmt"
	"log"
	"encoding/json"
	"github.com/wolf1996/gateway/token"
	"github.com/golang/protobuf/proto"
)

var ch *amqp.Channel
var q amqp.Queue

type QConfig struct {
	RabbitAddres string
	User string
	Pass string
}

func handler(msg resources.MessageTokened, tn token.Token)(err error){
	body, err := json.Marshal(msg)
	if err != nil {
		return
	}
	cds := amqp.Table{}
	mg, err :=proto.MarshalMessageSetJSON(tn)
	if err != nil {
		return
	}
	cds["token"] = mg

	mesg := amqp.Publishing{
		ContentType: "application/json",
		Body: body,
		Headers: cds,
	}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		mesg,
	)
	if err != nil {
		return
	}
	return
}

func DecrementEventsCounter(eventId int64, token token.Token) (err error){
	return handler(resources.MessageTokened{
		Token:  "",
		Message:DecrementRegistration{
			EventId: eventId,
		},
	},
	token)
}

func ApplyConfig(conf QConfig)(err error){
	log.Printf("User: %s, Addres: %s", conf.User, conf.RabbitAddres)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/",conf.User, conf.Pass, conf.RabbitAddres))
	if err != nil {
		log.Fatal(err.Error())
	}
	ch, err = conn.Channel()
	q,err = ch.QueueDeclare(
		"EventUsersDecrementMessages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("EventProducerError: %s",err.Error())
		return
	}
	return
}
