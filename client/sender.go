package client

import (
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"github.com/wolf1996/stats/shared"
	"github.com/wolf1996/stats/client/storage"
	"time"
)

type SenderConfig struct {
	BufferSize int
	Rabbit RabbitConfig
}

type MessageContainer struct {
	Mid shared.MessageId
	Data []byte
}

func StartSender(config SenderConfig, timeout int64)(err error, input chan MessageContainer){
	log.Printf("User: %s, Addres: %s", config.Rabbit.User, config.Rabbit.Addres)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/",config.Rabbit.User, config.Rabbit.Pass, config.Rabbit.Addres))
	if err != nil {
		log.Fatal(err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Responser Channel Creation Error: %s",err.Error())
		return
	}
	q, err := ch.QueueDeclare(
		shared.StatisticExchangeName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("Responser Channel Creation Error: %s",err.Error())
		return
	}
	input = make(chan MessageContainer, config.BufferSize)
	go sender(input, time.NewTicker(time.Duration(timeout)*time.Second).C, ch, q)
	return
}

func sender(input chan MessageContainer, ticker <-chan time.Time, ch *amqp.Channel, q amqp.Queue) {
	pblsh := func(container MessageContainer) {
		mesg := amqp.Publishing{
			ContentType: "application/json",
			Body: container.Data,
			Type: container.Mid.MsgType,
		}
		log.Printf("Send Message withId %v", container)
		err := ch.Publish(
			"",
			q.Name,
			false,
			false,
			mesg,
		)
		if err != nil {
			log.Printf("error in senders %s", err )
			return
		}
		datastring := string(container.Data[:])
		log.Printf("datastring is %s", datastring)
		storage.AddMessageStorage(container.Mid, datastring, container.Mid.MsgType)
	}

	refresh := func (){
		log.Printf("Start resending")
		msgs, err  := storage.GetAllMessages()
		if err != nil {
			log.Printf("Failing while getting messages %s", err.Error())
		}
		for id, data := range msgs{
			mid, err := storage.ParseKey(id)
			if err != nil {
				log.Printf("Failed To parse key %s", err.Error())
			}
			msgbdy, _ := storage.ParseVal(data)
			log.Printf("Msg Body is %s", msgbdy)
			container := MessageContainer{mid, []byte(msgbdy)}
			mesg := amqp.Publishing{
				ContentType: "application/json",
				Body: container.Data,
				Type: container.Mid.MsgType,
			}
			log.Printf("ReSend Message withId %v,\n msg = %s \n key = %s \n data = %s ", container,msgbdy,id,data)
			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				mesg,
			)
			if err != nil {
				log.Printf("Resend error get %s", err )
			}
		}
		log.Printf("Stop resending")
	}
	for {
		select {
		case msg := <-input:
			pblsh(msg)
		case <-ticker:
			refresh()
		}
	}
}