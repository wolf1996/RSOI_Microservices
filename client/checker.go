package client

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/wolf1996/stats/shared"
	"fmt"
	"github.com/wolf1996/stats/client/storage"
	"encoding/json"
)

type CheckerConfig struct {
	BufferSize int
	Rabbit RabbitConfig
}

func StartChecker(config CheckerConfig)(err error){
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
	err = ch.ExchangeDeclare(
		shared.ResponceExchangeName, // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Printf("Responser Channel Creation Error: %s",err.Error())
		return
	}
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Printf("Responser Que Creation Error: %s",err.Error())
		return
	}

	err = ch.QueueBind(
		q.Name,        // queue name
		producerName,             // routing key
		shared.ResponceExchangeName, // exchange
		false,
		nil)
	if err != nil {
		log.Printf("Responser Que Creation Error: %s",err.Error())
		return
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		log.Printf("Responser Que Creation Error: %s",err.Error())
		return
	}
	go checker(msgs)
	return
}


func checker(input <-chan amqp.Delivery) {
	check := func(delivery amqp.Delivery) {
		log.Printf("Got Message Response %s", string(delivery.Body[:]))
		var rsp shared.ResponceMsg
		err := json.Unmarshal(delivery.Body, &rsp)
		if err != nil {
			log.Printf("Failed to parse %s", err.Error())
			return
		}
		log.Printf("response serialized to %v", rsp)
		if rsp.ResponceStatus != 0 {
			log.Printf("Broken Message, refused")
		}
		err = storage.RemoveMessage(rsp.MessageIds)
		if err != nil {
			log.Printf("smth whent wrong with redis %s", err.Error())
		} else {
			log.Printf("Message removed from waiting list")
		}
	}
	for {
		select {
		case msg := <-input:
			check(msg)
		}
	}
}