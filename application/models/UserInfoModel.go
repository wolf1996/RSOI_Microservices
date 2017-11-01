package models

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

type DatabaseConfig struct {
	Username       string
	Pass           string
	DatabaseName   string
	DatabaseAddres string
}


var db *sql.DB

type UserInfo struct {
	Id 	  int64
	Name string
	Count int64
}

func PrepareTests()(err error){
	rows, err := db.Query("DROP TABLE IF EXISTS USER_INFO;"+
	"CREATE TABLE USER_INFO ("+
	"	ID       SERIAL PRIMARY KEY ,"+
	"	USERNAME VARCHAR(256),"+
	"	EVENTS_NUMBER INT DEFAULT(0)"+
	");"+
	"INSERT INTO USER_INFO VALUES  (DEFAULT, 'simpleUser', DEFAULT);)")
	if err!=nil {
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	return
}

func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres,config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}


func IncrementUserEventCounter(id string) (inf UserInfo, err error) {
	rows, err := db.Query("UPDATE USER_INFO SET EVENTS_NUMBER = EVENTS_NUMBER + 1 WHERE username = $1 RETURNING *;", id)
	if err!=nil {
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого ивента %d",id)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&inf.Id , &inf.Name, &inf.Count)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}

func GetUserInfo(login string) (inf UserInfo, err error) {
	rows, err := db.Query("SELECT * FROM USER_INFO WHERE username = $1", login)
	if err!=nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого пользователя %s",login)
		log.Print(err.Error())
		return
	}
	err =  rows.Scan(&inf.Id , &inf.Name, &inf.Count)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}


func DecrementUserEventCounter(id string) (inf UserInfo, err error) {
	rows, err := db.Query("UPDATE USER_INFO SET EVENTS_NUMBER = (CASE WHEN EVENTS_NUMBER > 0 THEN (EVENTS_NUMBER - 1) ELSE 0 END) WHERE username = $1 RETURNING *;", id)
	if err!=nil {
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого ивента %d",id)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&inf.Id , &inf.Name, &inf.Count)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}
