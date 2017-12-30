package worker

import (
"github.com/streadway/amqp"
"log"
"fmt"
"reflect"
"encoding/json"
	"github.com/wolf1996/stats/shared"
)


type HandlerConfig struct {
	RabConf 	    RabbitConfig
	NumberOfWorkers int
	BufferSize      int
}

func startHandlerWorkers(config HandlerConfig, output chan MessageContainer) (err error)  {
	log.Printf("Handler: %s, Addres: %s", config.RabConf.User, config.RabConf.Addres)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/",config.RabConf.User, config.RabConf.Pass, config.RabConf.Addres))
	if err != nil {
		log.Fatal(err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Handler Channel Creation Error: %s",err.Error())
		return
	}
	q, err := ch.QueueDeclare(
		shared.StatisticExchangeName,    // name
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("Handler Channel Creation Error: %s",err.Error())
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
	var inters  []chan interface{}
	var fins []chan error
	var finsCases []reflect.SelectCase
	for i:=0 ; i < config.NumberOfWorkers; i++ {
		inter, fin := startHandler(msgs, output)
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

func startHandler(msgs <-chan amqp.Delivery, output chan MessageContainer)(interr chan interface{}, finish chan error){
	interr = make(chan interface{})
	finish = make(chan error)
	go handler(msgs, output, interr, finish)
	return
}

func handler(msgs <-chan amqp.Delivery,
	outputs chan MessageContainer,
	interrupt <- chan interface{},
	finish chan error,){
	var err error
	defer func(){
		if err != nil {
			finish <- fmt.Errorf("Handle Worker stoped  with %v:",err)
		} else {
			finish <- fmt.Errorf("Handle Worker stoped without errors")
		}
	}()
	handle := func (msg amqp.Delivery)(errb error){
		log.Printf("Start Handler processing  %v", string(msg.Body[:len(msg.Body)]))
		tp := msg.Type
		switch(tp){
		case shared.TypeGetInf : {
			var bd shared.InfoViewMsg
			err := json.Unmarshal(msg.Body, &bd)
			if err != nil {
				log.Printf("Some shit in queue on Unmarshal")
				return
			}
			log.Printf("Send it to database")
			outputs <- MessageContainer{bd.Id, bd}
		}
		case shared.TypeChangeInf : {
			var bd shared.InfoChangeMsg
			err := json.Unmarshal(msg.Body, &bd)
			if err != nil {
				log.Printf("Some shit in queue on Unmarshal")
				return
			}
			log.Printf("Send it to database")
			outputs <- MessageContainer{bd.Id, bd}
		}

		case shared.TypeLogin : {
			var bd shared.LoginMsg
			err := json.Unmarshal(msg.Body, &bd)
			if err != nil {
				log.Printf("Some shit in queue on Unmarshal")
				return
			}
			log.Printf("Send it to database")
			outputs <- MessageContainer{bd.Id, bd}
		}
		default:
			log.Printf("Some shit in queue on type")
		}
		return
	}
WORKER_LOOP:
	for {
		select {
		case resp := <-msgs:
			err = handle(resp)
			if err != nil {
				break WORKER_LOOP
			}
		case <-interrupt:
			break WORKER_LOOP

		}
	}
}
