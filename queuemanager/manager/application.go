package manager

import (
	"github.com/streadway/amqp"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"log"
	"fmt"
	"encoding/json"
	"github.com/wolf1996/gateway/resources/userclient"
)

type RabbitConfig struct {
	Addres string
	User string
	Pass string
} 

type Config struct {
	Rabbit       RabbitConfig
	EventsConfig eventsclient.Config
	UsersConfig  userclient.Config
}

func UserIteration(msg amqp.Delivery)(err error){
	message := userclient.MessageTokened{}
	err = json.Unmarshal(msg.Body,&message)
	if err != nil {
		return
	}
	id := message.Message.UserId
	log.Printf("Processing User %d", id)
	_, err = userclient.DecrementEventsCounter(id)
	return
}

func EventIteration( msg amqp.Delivery)(err error){
	message := eventsclient.MessageTokened{}
	err = json.Unmarshal(msg.Body,&message)
	if err != nil {
		return
	}
	id := message.Message.EventId
	log.Printf("Processing Event %d", id)
	_, err = eventsclient.DecrementEventUsers(id)
	return
}

func UserLooperStarter(connection *amqp.Connection)( stop chan struct{}, finish chan error){
	stop = make(chan struct{})
	finish = make(chan error)
	go UserLooper(connection,stop,finish)
	return
}

func UserLooper(connection *amqp.Connection, stop chan struct{}, finish chan error)(err error){
	ch, err := connection.Channel()
	if err != nil {
		return
	}
	q,err := ch.QueueDeclare(
		"UserDecrementMessages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	MAINLOOP:
	for msg := range(msgs){
		err = UserIteration(msg)
		if err != nil {
			msg.Reject(true)
			log.Print(err.Error())
		} else {
			msg.Ack(false)
		}
		select{
		case <- finish:
			break MAINLOOP
		default:
		}
	}
	if err == nil {
		finish <- fmt.Errorf("finished")
	} else {
		finish <- err
	}
	return
}

func EventLooperStarter(connection *amqp.Connection)( stop chan struct{}, finish chan error){
	stop = make(chan struct{})
	finish = make(chan error)
	go EventLooper(connection,stop,finish)
	return
}

func EventLooper(connection *amqp.Connection,  stop chan struct{}, finish chan error){
	ch, err := connection.Channel()
	if err != nil {
		return
	}
	q,err := ch.QueueDeclare(
		"EventUsersDecrementMessages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	MAINLOOP:
	for msg := range(msgs){
		err = EventIteration(msg)
		if err != nil {
			msg.Reject(true)
			log.Print(err.Error())
		} else {
			msg.Ack(false)
		}
		select{
		case <- finish:
			break MAINLOOP
		default:
		}
	}
	if err == nil {
		finish <- fmt.Errorf("finished")
	} else {
		finish <- err
	}
	return
}

func initRabbit(conf RabbitConfig){
	log.Printf("Using rabbit: %s", conf.Addres)
	log.Printf("Starting rabbitmq user: %s", conf.User)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/",conf.User, conf.Pass, conf.Addres))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	stopEvent, finishEvent := EventLooperStarter(conn)
	stopUser, finishUser := UserLooperStarter(conn)
	select {
	case err = <-finishEvent:
		stopUser <- struct{}{}

	case err = <-finishUser:
		stopEvent <- struct{}{}
	}

}

func StartApplication(conf Config){
	eventsclient.SetConfigs(conf.EventsConfig)
	userclient.SetConfigs(conf.UsersConfig)
	go initRabbit(conf.Rabbit)
}