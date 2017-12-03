package userclient

import (
"github.com/wolf1996/gateway/resources"
"github.com/streadway/amqp"
"fmt"
"log"
"encoding/json"
)

var ch *amqp.Channel
var q amqp.Queue

type QConfig struct {
	RabbitAddres string
	User string
	Pass string
}

func handler(msg resources.MessageTokened)(err error){
	body, err := json.Marshal(msg)
	if err != nil {
		return
	}
	mesg := amqp.Publishing{
		ContentType: "application/json",
		Body: body,
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

func UserEventsDecrementCounter(userId string) (err error){
	return handler(resources.MessageTokened{
		Token:  "",
		Message:UserDecrementMessage{
			UserId: userId,
		},
	})
}

func ApplyConfig(conf QConfig)(err error){
	log.Printf("User: %s, Addres: %s", conf.User, conf.RabbitAddres)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/",conf.User, conf.Pass, conf.RabbitAddres))
	if err != nil {
		log.Fatal(err.Error())
	}
	ch, err = conn.Channel()
	q,err = ch.QueueDeclare(
		"UserDecrementMessages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("UserProducerError: %s",err.Error())
		return
	}
	return
}

