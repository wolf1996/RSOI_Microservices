package model

import (
	"gopkg.in/mgo.v2"
	"log"
	"github.com/wolf1996/stats/shared"
)

type MongoConfig struct {
	Addres string
	DbName string
}
var(
	mgoDb *mgo.Database
)

func ApplyConfig(config MongoConfig) (err error) {
	log.Printf("Mongo database addres userd %v, database is %v", config.Addres, config.DbName)
	session, err := mgo.Dial(config.Addres)
	if err != nil {
		return
	}
	session.Ping()
	if err != nil {
		return
	}
	session.SetSafe(&mgo.Safe{})
	mgoDb = session.DB(config.DbName)
	return
}

func GetLogins() (msgs []shared.LoginMsg, err error){
	log.Printf("Logins loaded")
	err = mgoDb.C(shared.TypeLogin).Find(nil).All(&msgs)
	log.Printf("Logins %v", msgs)

	return
}

func GetViewEvents() (msgs []shared.InfoViewMsg, err error) {
	log.Printf("Viewes loaded")
	err = mgoDb.C(shared.TypeGetInf).Find(nil).All(&msgs)
	log.Printf("Views %v", msgs)

	return
}


func GetChangeEvents() (msgs []shared.InfoChangeMsg, err error) {
	log.Printf("Changes loaded")
	err = mgoDb.C(shared.TypeChangeInf).Find(nil).All(&msgs)
	log.Printf("Changes %v", msgs)
	return
}