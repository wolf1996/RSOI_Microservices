package manager

import (
	"github.com/streadway/amqp"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"log"
	"fmt"
	"encoding/json"
	"github.com/wolf1996/gateway/resources"
)

type RabbitConfig struct {
	Addres string
	User string
	Pass string
} 

type Config struct {
	Rabbit       RabbitConfig
	EventsConfig eventsclient.Config
}

//func UserIteration(queue amqp.Queue, msg string)(err error){
//
//}

func EventIteration( msg amqp.Delivery)(err error){
	message := resources.MessageTokened{}
	err = json.Unmarshal(msg.Body,&message)
	if err != nil {
		return
	}
	reqData, ok  := message.Message.(eventsclient.DecrementRegistration)
	if !ok {
		err = fmt.Errorf("Wrong message type")
	}
	id := reqData.EventId
	_, err = eventsclient.DecrementEventUsers(id)
	return
}
//
//func UserLooperStarter(connection *amqp.Connection)( stop chan struct{}, finish chan error){
//	stop = make(chan struct{})
//	finish = make(chan error)
//	go UserLooper(connection,stop,finish)
//	return
//}
//
//func UserLooper(connection *amqp.Connection, stop chan struct{}, finish chan error)(err error){
//	ch, err := connection.Channel()
//	if err != nil {
//		return
//	}
//	q,err := ch.QueueDeclare(
//		"UserDecrementMessages",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		return
//	}
//	msgs, err := ch.Consume(
//		q.Name,
//		"",
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		return
//	}
//	MAINLOOP:
//	for msg := range(msgs){
//		UserIteration(q,string(msg.Body[:]))
//		if err != nil {
//			break MAINLOOP
//		}
//		select{
//		case <- finish:
//			break MAINLOOP
//		default:
//		}
//	}
//	if err == nil {
//		finish <- fmt.Errorf("finished")
//	} else {
//		finish <- err
//	}
//	return
//}

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
		true,
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
			log.Print(err.Error())
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
	_, finishEvent := EventLooperStarter(conn)
	//stopUser, finishUser := UserLooperStarter(conn)
	//select {
	//case err = <-finishEvent:
	//	stopUser <- struct{}{}
	//
	//case err = <-finishUser:
	//	stopEvent <- struct{}{}
	//}
	select{
	case err:= <-finishEvent:
		log.Print(err.Error())
	}
}

func StartApplication(conf Config){
	go initRabbit(conf.Rabbit)
	eventsclient.SetConfigs(conf.EventsConfig)
}