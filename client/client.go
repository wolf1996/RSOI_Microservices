package client

import (
	"log"
	"github.com/wolf1996/stats/shared"
	"time"
	"math/rand"
	"encoding/json"
	"github.com/wolf1996/stats/client/storage"
)

type Config struct {
	ProducerName string
	TimeOut      int64
	Sconfig  SenderConfig
	ChConfig CheckerConfig
	RedConf  storage.Config
}

type RabbitConfig struct {
	Addres string
	User string
	Pass string
}

var (
	producerName string
	input chan MessageContainer
)

func StartApplication(config Config) (err error) {
	producerName = config.ProducerName
	log.Printf("statistcs with name %s, timeout %d", producerName, config.TimeOut)
	err = storage.ApplyConfig(config.RedConf)
	if err != nil {
		return
	}
	err = StartChecker(config.ChConfig)
	if err != nil {
		return
	}
	err, input = StartSender(config.Sconfig, config.TimeOut)
	if err != nil {
		close(input)
		return
	}
	return
}

func WriteInfoViewMessage(path, userId string)(err error){
	msg := shared.InfoViewMsg{
		Id:shared.MessageId{
			Producer:producerName,
			Timestamp:time.Now().Unix(),
			Random:rand.Int(),
			MsgType:shared.TypeGetInf,
		},
		Path:path,
		UserId:userId,
	}
	btmsg, err := json.Marshal(msg)
	if err != nil {
		return
	}
	cont := MessageContainer{Mid:msg.Id, Data:btmsg}
	input <- cont
	return
}

func WriteInfoChangeMessage(path, userId string)(err error){
	msg := shared.InfoChangeMsg{
		Id:shared.MessageId{
			Producer:producerName,
			Timestamp:time.Now().Unix(),
			Random:rand.Int(),
			MsgType:shared.TypeChangeInf,
		},
		Path:path,
		UserId:userId,
	}
	btmsg, err := json.Marshal(msg)
	if err != nil {
		return
	}
	cont := MessageContainer{Mid:msg.Id, Data:btmsg}
	input <- cont
	return
}

func WriteLoginMessage( okLog bool, info string)(err error){
	msg := shared.LoginMsg{
		Id:shared.MessageId{
			Producer:producerName,
			Timestamp:time.Now().Unix(),
			Random:rand.Int(),
			MsgType:shared.TypeLogin,
		},
		Ok: okLog,
		Info: info,
	}
	btmsg, err := json.Marshal(msg)
	if err != nil {
		return
	}
	cont := MessageContainer{Mid:msg.Id, Data:btmsg}
	input <- cont
	return
}