package worker

import (
	"log"
	"gopkg.in/mgo.v2"
	"reflect"
	"fmt"
	"github.com/wolf1996/stats/shared"
)

type DatabaseWorkerConfig struct {
	BufferSize    int
	NumberOfWorkers    int
	Types []string
	MgoConf MongoConfig
}

type MongoConfig struct {
	Addres string
	DbName string
}

type MessageContainer struct {
	Id shared.MessageId
	Data interface{}
}

func StartDatabaseWorker(config DatabaseWorkerConfig, output chan shared.ResponceMsg) ( input chan MessageContainer, err error){
	log.Printf("Mongo database addres userd %v, database is %v", config.MgoConf.Addres, config.MgoConf.DbName)
	session, err := mgo.Dial(config.MgoConf.Addres)
	if err != nil {
		return
	}
	session.Ping()
	if err != nil {
		return
	}
	session.SetSafe(&mgo.Safe{})
	db := session.DB(config.MgoConf.DbName)
	index := mgo.Index{
		Key: []string{"message_ids.timestamp", "message_ids.producer","message_ids.random"},
		Unique:true,
	}
	for _, i := range config.Types {
		err = db.C(i).EnsureIndex(index)
		if err != nil {
			return
		}
	}
	input = make(chan MessageContainer, config.BufferSize)
	var inters  []chan interface{}
	var fins []chan error
	var finsCases []reflect.SelectCase
	for i:=0 ; i < config.NumberOfWorkers; i++ {
		inter, fin := startDatabaser(db,config.Types, input, output)
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
		session.Close()
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

func startDatabaser(database *mgo.Database,types []string, input chan MessageContainer,output chan shared.ResponceMsg) (interr chan interface{}, finish chan error){
	interr = make(chan interface{})
	finish = make(chan error)
	go databaser(database,types, input, output, interr, finish)
	return
}

func databaser( database *mgo.Database,types []string, input chan MessageContainer,output chan shared.ResponceMsg, interr chan interface{}, finish chan error) {
	var err error
	defer func(){
		if err != nil {
			finish <- fmt.Errorf("Worker stoped  with %v:",err)
		} else {
			finish <- fmt.Errorf("Worker stoped without errors")
		}
	}()

	writer := func (msg MessageContainer)(err error){
		log.Printf("Start Database processing  %v", msg)
		exists := false
		for _, i := range types{
			if i == msg.Id.MsgType {
				exists = true
			}
		}
		if !exists{
			log.Printf("Message Rejected %v", msg)
			output <- shared.ResponceMsg{1, msg.Id}
			return
		}
		collection := database.C(msg.Id.MsgType)
		err = collection.Insert(msg.Data)
		if err != nil {
			log.Printf(err.Error())
			log.Printf("Message Rejected by database %v", msg)
			output <- shared.ResponceMsg{1,msg.Id}
			if mgo.IsDup(err){
				err = nil
			}
		} else {
			log.Printf("Message inserted  %v", msg)
			output <-shared.ResponceMsg{0,msg.Id}
		}
		return
	}
WORKER_LOOP:
	for {
		select {
		case msg := <-input:
			err = writer(msg)
			if err != nil {
				break WORKER_LOOP
			}
		case <-interr:
			break WORKER_LOOP

		}
	}
}