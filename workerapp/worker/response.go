package worker

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
	"reflect"
	"encoding/json"
	"github.com/wolf1996/stats/shared"
)

type ResponserConfig struct {
	RabConf 	    RabbitConfig
	NumberOfWorkers int
	BufferSize      int
}

func startResponserWorkers(config ResponserConfig) (input chan shared.ResponceMsg,err error)  {
	log.Printf("Responser: %s, Addres: %s", config.RabConf.User, config.RabConf.Addres)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/",config.RabConf.User, config.RabConf.Pass, config.RabConf.Addres))
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
	input = make(chan shared.ResponceMsg, config.BufferSize)
	var inters  []chan interface{}
	var fins []chan error
	var finsCases []reflect.SelectCase
	for i:=0 ; i < config.NumberOfWorkers; i++ {
		inter, fin := startResponser(ch, input)
		finsCases = append(finsCases, reflect.SelectCase{Dir:reflect.SelectRecv, Chan: reflect.ValueOf(fin)})
		fins = append(fins, fin)
		inters = append(inters, inter)
	}

	go func() {
		chsn, val, _ := reflect.Select(finsCases)
		log.Printf("Worker Crashed Kill'm all %v", val)
		fins[chsn] = fins[len(fins)-1]
		fins = fins[:len(fins)-1]
		for _, chnl := range inters{
			chnl <- struct{}{}
		}
		for i, chnl := range fins{
			<-chnl
			log.Print("Worker %d Stoped",i )
		}
		log.Print("Workers Stoped")
		ch.Close()
		for _, chnl := range fins{
			close(chnl)
		}
		for _, chnl := range inters{
			close(chnl)
		}
		log.Print("Stoped, quit")
	}()
	return
}

func startResponser(channel *amqp.Channel, input <-chan shared.ResponceMsg)(interr chan interface{}, finish chan error){
	interr = make(chan interface{})
	finish = make(chan error)
	go responcer(channel, input, interr, finish)
	return
}

func responcer( channel *amqp.Channel,
	input <-chan shared.ResponceMsg,
	interrupt <- chan interface{},
	finish chan error,){
	var err error
	defer func(){
		if err != nil {
			finish <- fmt.Errorf("Worker stoped  with %v:",err)
		} else {
			finish <- fmt.Errorf("Worker stoped without errors")
		}
	}()
	sender := func (msg shared.ResponceMsg)(err error){
		log.Printf("Start Responce processing  %v", msg)
		body, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Marshal error while processing  %v", msg)
			return
		}
		mesg := amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		}
		err = channel.Publish(
			shared.ResponceExchangeName,
			msg.MessageIds.Producer,
			false,
			false,
			mesg,
		)
		return
	}
WORKER_LOOP:
	for {
		select {
		case resp := <-input:
			err = sender(resp)
			if err != nil {
				break WORKER_LOOP
			}
		case <-interrupt:
			break WORKER_LOOP

		}
	}
}